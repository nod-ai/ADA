#!/usr/bin/env python3
# Copyright (C) 2025 Advanced Micro Devices, Inc.
# All rights reserved.

# If the requests module is not installed on your system please run the following commands:
# python3 -m pip install --upgrade pip
# python3 -m pip install requests

import sys
import os
import re
import time
import json
import argparse
import requests
from datetime import datetime
from requests.adapters import HTTPAdapter
from requests.packages.urllib3.util.retry import Retry

# Suppress InsecureRequestWarning if using verify=False in requests
import urllib3

urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)

# --------------------------------------------------------------------
# Variables for porting between platforms
# --------------------------------------------------------------------
REDFISH_OEM = "redfish/v1/Oem/Supermicro/MI300X"
REDFISH_MANAGER_BMC = "redfish/v1/Managers/1"
REDFISH_SYSTEM_BMC = "redfish/v1/Systems/1"

# --------------------------------------------------------------------
# Setup argument parsing for optional debug mode
# --------------------------------------------------------------------
parser = argparse.ArgumentParser(description="Erase all MI300X logs via Redfish")
parser.add_argument("--debug", action="store_true", help="Enable debug output")
args = parser.parse_args()
DEBUG = args.debug

# --------------------------------------------------------------------
# Helper Functions
# --------------------------------------------------------------------
def log(message: str):
    """Print a timestamped log message."""
    print(f"{datetime.now().strftime('%Y-%m-%d %H:%M:%S')} - {message}")


def check_response_success(response, error_message: str):
    """Check if a requests.Response object is successful, exit otherwise."""
    if not response.ok:
        log(error_message)
        log(f"HTTP status code: {response.status_code}")
        sys.exit(1)


def prompt_for_bmc_credentials():
    """Prompt for BMC Username and Password if they are not set in environment variables."""
    bmc_username = os.environ.get("BMC_USERNAME")
    bmc_password = os.environ.get("BMC_PASSWORD")

    if not bmc_username:
        bmc_username = input("Enter BMC Username: ")
    if not bmc_password:
        if DEBUG:
            bmc_password = input("Enter BMC Password: ")
        else:
            import getpass

            bmc_password = getpass.getpass("Enter BMC Password: ")

    return bmc_username, bmc_password


def prompt_for_bmc_ip():
    """Prompt for BMC IP address if not set and validate the format."""
    bmc_ip = os.environ.get("BMC_IP", "")
    ip_pattern = re.compile(r"^\d+\.\d+\.\d+\.\d+$")

    while not ip_pattern.match(bmc_ip):
        bmc_ip = input("Enter valid BMC IP Address: ")

    return bmc_ip


def http_request_with_retries(method, url, max_retries=3, delay=2, **kwargs):
    """
    Attempt an HTTP request up to `max_retries` times. If a transient error
    like a connection timeout or reset occurs, wait `delay` seconds before
    trying again.

    :param method: "get", "post", etc.
    :param url: The URL to request.
    :param max_retries: Maximum number of total attempts.
    :param delay: Delay in seconds between retry attempts.
    :param kwargs: Additional arguments to pass to requests.request().
    :return: requests.Response object on success; sys.exit on repeated failures.
    """

    """Create a requests session with retry logic."""
    session = requests.Session()
    retries = Retry(
        total=5,  # Number of retries
        backoff_factor=2,  # Exponential backoff factor
        status_forcelist=[500, 502, 503, 504],  # Retry on these HTTP errors
        allowed_methods=["GET", "POST"],
    )
    session.mount("https://", HTTPAdapter(max_retries))

    for attempt in range(1, max_retries + 1):
        try:
            if method == "get":
                response = session.get(url, verify=False, timeout=10, **kwargs)
            elif method == "post":
                response = session.post(url, verify=False, timeout=10, **kwargs)
            else:
                raise ValueError("Unsupported HTTP method")
            return response
        except (
            requests.exceptions.ConnectionError,
            requests.exceptions.Timeout,
            requests.exceptions.RequestException,
        ) as e:
            # These are typical transient errors that can occur on unreliable networks
            if attempt < max_retries:
                log(
                    f"Request error: {e}. Retrying in {delay} seconds (attempt {attempt}/{max_retries})."
                )
                time.sleep(delay)
            else:
                log(f"Request failed after {max_retries} attempts: {e}")
                sys.exit(1)

    # If somehow we got here, return None or exit
    # (normally you'd never get here because of sys.exit).
    sys.exit(1)


def authenticate(bmc_ip, bmc_username, bmc_password):
    """
    Check if the provided IP, username, and password are valid by querying
    the Redfish Managers resource.
    """
    url = f"https://{bmc_ip}/{REDFISH_MANAGER_BMC}"
    try:
        response = http_request_with_retries(
            "get", url, auth=(bmc_username, bmc_password)
        )
        if not response.ok:
            log("Authentication check failed. HTTP code: " + str(response.status_code))
            return False
        data = response.json()
        # If there's an "error" key in the response, treat it as auth failure
        if "error" in data:
            return False
    except Exception as e:
        log(f"Exception while authenticating: {e}")
        return False

    return True


def tasks_wait(bmc_ip, bmc_username, bmc_password):
    """
    Wait for all tasks in the TaskService to complete. Exits the script
    if any task fails or times out.
    """
    TIMEOUT_SECONDS = 25 * 60  # 25 minutes
    INTERVAL = 3

    # Check how many tasks are there
    url = f"https://{bmc_ip}/{REDFISH_OEM}/TaskService/Tasks"
    try:
        response = http_request_with_retries(
            "get", url, auth=(bmc_username, bmc_password)
        )
        check_response_success(response, "Failed to query tasks.")
        data = response.json()
        tasks = data.get("Members@odata.count", 0)
    except Exception as e:
        log(f"Exception while querying tasks: {e}")
        sys.exit(1)

    if DEBUG:
        log(f"Tasks: {tasks}")

    # If no tasks, nothing to wait for
    if tasks == 0:
        return

    for task_index in range(tasks):
        newline_needed = False
        elapsed_time = 0

        while True:
            # Query the status of each task
            task_url = f"https://{bmc_ip}/{REDFISH_OEM}/TaskService/Tasks/{task_index}"
            try:
                response = http_request_with_retries(
                    "get", task_url, auth=(bmc_username, bmc_password)
                )
                check_response_success(
                    response, f"Failed to query status of task {task_index}."
                )
                status_data = response.json()
                status = status_data.get("TaskState", "Unknown")
            except Exception as e:
                log(f"Exception while querying task {task_index}: {e}")
                sys.exit(1)

            if status == "Completed":
                if DEBUG:
                    print(f"\nTask {task_index} completed successfully.")
                break
            elif status == "Failed":
                log(f"Task {task_index} failed.")
                sys.exit(1)
            elif status in ["Running", "New", "Pending"]:
                newline_needed = True
                print(
                    f"\rTask {task_index} still running, elapsed time {elapsed_time}s",
                    end="",
                )
                sys.stdout.flush()
                time.sleep(INTERVAL)
                elapsed_time += INTERVAL
            else:
                newline_needed = True
                print(
                    f"\rUnknown task status: {status}, task {task_index}, elapsed time {elapsed_time}s",
                    end="",
                )
                sys.stdout.flush()
                time.sleep(INTERVAL)
                elapsed_time += INTERVAL

            if elapsed_time > TIMEOUT_SECONDS:
                log(
                    f"Task {task_index} failed to complete in {TIMEOUT_SECONDS // 60} minutes."
                )
                sys.exit(1)

        if newline_needed:
            print()  # Move to the next line


def systemPowerOn(bmc_ip, bmc_username, bmc_password):
    """
    Power on system using BMC Redfish, then wait for host to come up.
    """
    dwell = 20
    INTERVAL = 3
    TIMEOUT_SECONDS = 8 * 60

    log("Powering system on")

    reset_url = f"https://{bmc_ip}/{REDFISH_SYSTEM_BMC}/Actions/ComputerSystem.Reset"
    data = {"Action": "Reset", "ResetType": "On"}

    try:
        response = http_request_with_retries(
            "post", reset_url, auth=(bmc_username, bmc_password), json=data
        )
        check_response_success(response, "Host power on failure.")
        task_response_text = response.text
    except Exception as e:
        log(f"Exception while Host power on: {e}")
        sys.exit(1)

    log(f"Waiting for {dwell} seconds to confirm power is on")
    time.sleep(dwell)

    bmc_url = f"https://{bmc_ip}/{REDFISH_SYSTEM_BMC}"

    try:
        response = http_request_with_retries(
            "get", bmc_url, auth=(bmc_username, bmc_password)
        )
        check_response_success(response, "Failed to get power state.")
        resp = response.json()
        powerState = resp.get("PowerState", "Unknown")
        if powerState != "On":
            log(f"Failed to power on system, power state is {powerState}")
            sys.exit(1)
        else:
            log(f"System is {powerState}")
    except Exception as e:
        log(f"Exception while powering on system: {e}")
        sys.exit(1)

    log(
        f"Sleep for {TIMEOUT_SECONDS // 60} minutes so we can be sure BMC and SMC are ready"
    )

    elapsed_time = 0

    while elapsed_time < TIMEOUT_SECONDS:
        print(f"\rElapsed time {elapsed_time}s", end="")
        sys.stdout.flush()
        time.sleep(INTERVAL)
        elapsed_time += INTERVAL

    print()  # Move to next line


def systemPowerOff(bmc_ip, bmc_username, bmc_password):
    """
    Power off system using BMC Redfish.
    Host must be up for ForceOff or GracefulShutdown to work reliably.
    """
    dwell = 20
    log("Powering system off")

    reset_url = f"https://{bmc_ip}/{REDFISH_SYSTEM_BMC}/Actions/ComputerSystem.Reset"
    data = {"Action": "Reset", "ResetType": "ForceOff"}

    try:
        response = http_request_with_retries(
            "post", reset_url, auth=(bmc_username, bmc_password), json=data
        )
        check_response_success(response, "Host power off failure.")
        task_response_text = response.text
    except Exception as e:
        log(f"Exception while Host power off: {e}")
        sys.exit(1)

    log(f"Waiting for {dwell} seconds to confirm power is off")
    time.sleep(dwell)

    bmc_url = f"https://{bmc_ip}/{REDFISH_SYSTEM_BMC}"

    try:
        response = http_request_with_retries(
            "get", bmc_url, auth=(bmc_username, bmc_password)
        )
        check_response_success(response, "Failed to get power state.")
        resp = response.json()
        powerState = resp.get("PowerState", "Unknown")
        if powerState == "On":
            log("Failed to power off system")
            sys.exit(1)
        else:
            log(f"System is {powerState}")
    except Exception as e:
        log(f"Exception while powering off system: {e}")
        sys.exit(1)


def systemPowerCycle(bmc_ip, bmc_username, bmc_password):
    """
    Cycle the power using the BMC.
    This implies an off-then-wait-then-on sequence.
    """
    dwell = 120
    log("Initiating DC power cycle")
    systemPowerOff(bmc_ip, bmc_username, bmc_password)
    log(f"Dropping power for {dwell // 60} minutes")
    time.sleep(dwell)
    systemPowerOn(bmc_ip, bmc_username, bmc_password)


def logsClear(bmc_ip, bmc_username, bmc_password):
    """
    Clear all logs on UBB.
    """
    log("Clearing all logs on UBB")

    paths = [
        f"https://{bmc_ip}/{REDFISH_OEM}/Systems/UBB/LogServices/Dump/Actions/LogService.ClearLog",
        f"https://{bmc_ip}/{REDFISH_OEM}/Systems/UBB/LogServices/EventLog/Actions/LogService.ClearLog",
        f"https://{bmc_ip}/{REDFISH_OEM}/Systems/UBB/LogServices/DiagLogs/Actions/LogService.ClearLog",
    ]

    for p in paths:
        try:
            payload = {}
            response = http_request_with_retries(
                "post", p, auth=(bmc_username, bmc_password), json=payload
            )
            check_response_success(response, "Failed to clear logs at {p}.")
            task_response_text = response.text
        except Exception as e:
            log(f"Exception while clearing log data: {e}")
            sys.exit(1)
        time.sleep(3)


def main():
    # ----------------------------------------------------------------
    # 1. Prompt for credentials & IP
    # ----------------------------------------------------------------
    bmc_username, bmc_password = prompt_for_bmc_credentials()
    bmc_ip = prompt_for_bmc_ip()

    # ----------------------------------------------------------------
    # 2. Test authentication
    # ----------------------------------------------------------------
    if not authenticate(bmc_ip, bmc_username, bmc_password):
        log("Authentication failed, check IP address, username, and password")
        log(f"IP address used: {bmc_ip}")
        log(f"Username used: {bmc_username}")
        if DEBUG:
            log(f"Password used: {bmc_password}")
        sys.exit(1)

    # ----------------------------------------------------------------
    # 3. Clear the logs
    # ----------------------------------------------------------------
    logsClear(bmc_ip, bmc_username, bmc_password)

    # ----------------------------------------------------------------
    # 4. Cycle power so that new logs generate
    # ----------------------------------------------------------------
    systemPowerCycle(bmc_ip, bmc_username, bmc_password)

    log(f"All done.")


if __name__ == "__main__":
    main()
