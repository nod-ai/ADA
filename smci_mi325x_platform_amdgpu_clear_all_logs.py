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
# Setup argument parsing for optional debug mode
# --------------------------------------------------------------------
parser = argparse.ArgumentParser(description="Erase all MI300X logs via Redfish")
parser.add_argument("--debug", action="store_true", help="Enable debug output")
parser.add_argument("--verbose", action="store_true", help="Enable verbose output")
args = parser.parse_args()
DEBUG = args.debug
VERBOSE = args.verbose

# --------------------------------------------------------------------
# Constants for porting between platforms
# --------------------------------------------------------------------
REDFISH_OEM = "redfish/v1"
REDFISH_MANAGER_BMC = "redfish/v1/Managers/1"
REDFISH_SYSTEM_BMC = "redfish/v1/Systems/1"
PORT = 443
PROTOCOL = "https"
POWER_ON = {"Action": "Reset", "ResetType": "On"}
POWER_OFF = {"Action": "Reset", "ResetType": "ForceOff"}

# --------------------------------------------------------------------
# Helper Functions
# --------------------------------------------------------------------
def log(message: str):
    """Print a timestamped log message."""
    print(f"{datetime.now().strftime('%Y-%m-%d %H:%M:%S')} - {message}")


def check_response_success(response, error_message: str):
    """
    Check if a requests.Response object is successful; if not, log a
    message with status code and exit. Specifically check for 403/404
    to give more clarity.
    """
    if response.status_code == 403:
        log("403 Forbidden: Possible username/password error.")
        log(f"HTTP status code: {response.status_code}")
        sys.exit(1)
    elif response.status_code == 404:
        log("404 Not Found: The specified resource or URL could not be found.")
        log(f"HTTP status code: {response.status_code}")
        sys.exit(1)

    if not response.ok:
        log(error_message)
        log(f"HTTP status code: {response.status_code}")
        sys.exit(1)


def prompt_for_bmc_credentials():
    """Prompt for BMC Username and Password if they are not set in environment variables."""
    bmc_username = os.environ.get("BMC_USERNAME")
    bmc_password = os.environ.get("BMC_PASSWORD")

    while not bmc_username:
        bmc_username = input("Enter BMC Username: ")
    while not bmc_password:
        if DEBUG:
            bmc_password = input("Enter BMC Password: ")
        else:
            import getpass

            bmc_password = getpass.getpass("Enter BMC Password: ")

    return bmc_username, bmc_password


def prompt_for_bmc_ip():
    """Prompt for BMC IP address if not set and validate the format."""
    bmc_ip = os.environ.get("BMC_IP", "")
    ip_pattern = re.compile(r"^\d+\.\d+\.\d+\.\d+$")  # Simple IPv4 check

    while not ip_pattern.match(bmc_ip):
        bmc_ip = input("Enter valid BMC IP Address (IPv4): ")

    return bmc_ip


def http_request_with_retries(method, url, max_retries=3, delay=2, **kwargs):
    """
    Attempt an HTTP request up to `max_retries` times. If a transient error
    (connection timeout, reset, etc.) occurs, wait `delay` seconds before
    trying again.

    :param method: "get", "post", etc.
    :param url: The URL to request.
    :param max_retries: Maximum number of total attempts.
    :param delay: Delay in seconds between retry attempts.
    :param kwargs: Additional arguments to pass to requests.request().
    :return: requests.Response object on success; sys.exit(1) on repeated failures.
    """
    session = requests.Session()
    retries = Retry(
        total=max_retries,  # number of retries
        backoff_factor=2,  # exponential backoff factor
        status_forcelist=[500, 502, 503, 504],  # retry only on these status codes
        allowed_methods=["GET", "POST", "PATCH"],
    )
    session.mount("https://", HTTPAdapter(max_retries=retries))
    session.mount("http://", HTTPAdapter(max_retries=retries))

    if VERBOSE:
        print(f"{method} method for URL: {url}")

    for attempt in range(1, max_retries + 1):
        try:
            if method.lower() == "get":
                response = session.get(url, verify=False, timeout=10, **kwargs)
            elif method.lower() == "post":
                response = session.post(url, verify=False, timeout=10, **kwargs)
            elif method.lower() == "patch":
                response = session.patch(url, verify=False, timeout=10, **kwargs)
            else:
                raise ValueError(f"Unsupported HTTP method: {method}")

            if VERBOSE:
                print(f"Response {attempt}: {response.text}")

            return response

        except (
            requests.exceptions.ConnectionError,
            requests.exceptions.Timeout,
            requests.exceptions.RequestException,
        ) as e:
            if attempt < max_retries:
                print(
                    f"Request error: {e}. Retrying in {delay} seconds "
                    f"(attempt {attempt}/{max_retries})."
                )
                time.sleep(delay)
            else:
                print(f"Request failed after {max_retries} attempts: {e}")
                sys.exit(1)

    # Should never reach here due to sys.exit(1)
    sys.exit(1)


def authenticate(bmc_ip, bmc_username, bmc_password):
    """
    Check if the provided IP, username, and password are valid by querying
    the Redfish Managers resource.
    """
    url = f"{PROTOCOL}://{bmc_ip}:{PORT}/{REDFISH_MANAGER_BMC}"
    try:
        response = http_request_with_retries(
            "get", url, auth=(bmc_username, bmc_password)
        )
        if not response.ok:
            log(f"Authentication check failed. HTTP code: {response.status_code}")
            return False
        data = response.json()
        if "error" in data:
            # Some BMCs return an 'error' key if authentication fails
            return False
    except Exception as e:
        log(f"Exception while authenticating: {e}")
        return False

    return True


def systemPowerOn(bmc_ip, bmc_username, bmc_password):
    """
    Power on system using BMC Redfish, then wait for host to come up.
    """
    dwell = 20
    INTERVAL = 3
    TIMEOUT_SECONDS = 8 * 60

    log("Powering system on")

    reset_url = f"{PROTOCOL}://{bmc_ip}:{PORT}/{REDFISH_SYSTEM_BMC}/Actions/ComputerSystem.Reset"

    try:
        response = http_request_with_retries(
            "post", reset_url, auth=(bmc_username, bmc_password), json=POWER_ON
        )
        check_response_success(response, "Host power on failure.")
        task_response_text = response.text
    except Exception as e:
        log(f"Exception while Host power on: {e}")
        sys.exit(1)

    log(f"Waiting for {dwell} seconds to confirm power is on")
    time.sleep(dwell)

    bmc_url = f"{PROTOCOL}://{bmc_ip}:{PORT}/{REDFISH_SYSTEM_BMC}"

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

    remaining_time = TIMEOUT_SECONDS

    while remaining_time > 0:
        print(f"\rRemaining time: {remaining_time:>3}s", end="")
        sys.stdout.flush()
        time.sleep(INTERVAL)
        remaining_time -= INTERVAL

    print()  # Move to next line


def systemPowerOff(bmc_ip, bmc_username, bmc_password):
    """
    Power off system using BMC Redfish.
    Host must be up for ForceOff or GracefulShutdown to work reliably.
    """
    dwell = 20
    log("Powering system off")

    reset_url = f"{PROTOCOL}://{bmc_ip}:{PORT}/{REDFISH_SYSTEM_BMC}/Actions/ComputerSystem.Reset"

    try:
        response = http_request_with_retries(
            "post", reset_url, auth=(bmc_username, bmc_password), json=POWER_OFF
        )
        check_response_success(response, "Host power off failure.")
        task_response_text = response.text
    except Exception as e:
        log(f"Exception while Host power off: {e}")
        sys.exit(1)

    log(f"Waiting for {dwell} seconds to confirm power is off")
    time.sleep(dwell)

    bmc_url = f"{PROTOCOL}://{bmc_ip}:{PORT}/{REDFISH_SYSTEM_BMC}"

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
        f"{PROTOCOL}://{bmc_ip}:{PORT}/{REDFISH_OEM}/Systems/UBB/LogServices/Dump/Actions/LogService.ClearLog",
        f"{PROTOCOL}://{bmc_ip}:{PORT}/{REDFISH_OEM}/Systems/UBB/LogServices/EventLog/Actions/LogService.ClearLog",
        f"{PROTOCOL}://{bmc_ip}:{PORT}/{REDFISH_OEM}/Systems/UBB/LogServices/DiagLogs/Actions/LogService.ClearLog",
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
