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
parser = argparse.ArgumentParser(description="Collect AllLogs from BMC via Redfish")
parser.add_argument("--debug", action="store_true", help="Enable debug output")
args = parser.parse_args()
DEBUG = args.debug

# --------------------------------------------------------------------
# Constants for porting between platforms
# --------------------------------------------------------------------
REDFISH_OEM = "redfish/v1/Oem/Supermicro/MI300X"
REDFISH_MANAGER_BMC = "redfish/v1/Managers/1"
PORT = 443
PROTOCOL = "https"

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


def tasks_wait(bmc_ip, bmc_username, bmc_password):
    """
    Wait for all tasks in the TaskService to complete. Exits the script
    if any task fails or if it times out.
    """
    TIMEOUT_SECONDS = 25 * 60  # 25 minutes
    INTERVAL = 3

    # Give the BMC a moment to register the newly-triggered tasks
    time.sleep(3)

    # Check how many tasks are have been created
    url = f"{PROTOCOL}://{bmc_ip}:{PORT}/{REDFISH_OEM}/TaskService/Tasks"
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
        log(f"Tasks found: {tasks}")

    # If no tasks, nothing to wait for
    if tasks == 0:
        return

    for task_index in range(tasks):
        newline_needed = False
        elapsed_time = 0

        while True:
            # Query the status of each task
            task_url = f"{PROTOCOL}://{bmc_ip}:{PORT}/{REDFISH_OEM}/TaskService/Tasks/{task_index}"
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
                # Show ongoing progress
                newline_needed = True
                print(
                    f"\rTask {task_index} is still running, elapsed {elapsed_time}s",
                    end="",
                )
                sys.stdout.flush()
                time.sleep(INTERVAL)
                elapsed_time += INTERVAL
            else:
                newline_needed = True
                print(
                    f"\rUnknown task status: {status}, task {task_index}, elapsed {elapsed_time}s",
                    end="",
                )
                sys.stdout.flush()
                time.sleep(INTERVAL)
                elapsed_time += INTERVAL

            if elapsed_time > TIMEOUT_SECONDS:
                log(
                    f"Task {task_index} failed to complete within {TIMEOUT_SECONDS // 60} minutes."
                )
                sys.exit(1)

        if newline_needed:
            print()  # End the progress line


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
        log("Authentication failed. Please check IP address, username, and password.")
        log(f"IP address used: {bmc_ip}")
        log(f"Username used: {bmc_username}")
        if DEBUG:
            log(f"Password used: {bmc_password}")
        sys.exit(1)

    # ----------------------------------------------------------------
    # 3. Collect diagnostic data (AllLogs)
    # ----------------------------------------------------------------
    log("Collecting AllLogs file...")
    collect_url = (
        f"{PROTOCOL}://{bmc_ip}:{PORT}/{REDFISH_OEM}/Systems/UBB/"
        "LogServices/DiagLogs/Actions/LogService.CollectDiagnosticData"
    )
    payload = {"DiagnosticDataType": "OEM", "OEMDiagnosticDataType": "AllLogs"}

    try:
        response = http_request_with_retries(
            "post", collect_url, auth=(bmc_username, bmc_password), json=payload
        )
        check_response_success(response, "Failed to collect diagnostic data.")
        task_response_text = response.text
    except Exception as e:
        log(f"Exception while collecting diagnostic data: {e}")
        sys.exit(1)

    # Extract Task ID from the response text
    match = re.search(r'Tasks/([^"]+)', task_response_text)
    if not match:
        log("Script failed, no valid task ID found in response.")
        sys.exit(1)
    tasks_id = match.group(1)  # Not used further, but confirms presence

    # ----------------------------------------------------------------
    # 4. Wait for tasks to complete
    # ----------------------------------------------------------------
    tasks_wait(bmc_ip, bmc_username, bmc_password)

    # ----------------------------------------------------------------
    # 5. Get count of diagnostic log entries
    # ----------------------------------------------------------------
    entries_url = f"{PROTOCOL}://{bmc_ip}:{PORT}/{REDFISH_OEM}/Systems/UBB/LogServices/DiagLogs/Entries"
    try:
        response = http_request_with_retries(
            "get", entries_url, auth=(bmc_username, bmc_password)
        )
        check_response_success(response, "Failed to get diagnostic log entries.")
        diag_data = response.json()
        entries_count = diag_data.get("Members@odata.count", 0)
    except Exception as e:
        log(f"Exception while retrieving log entries: {e}")
        sys.exit(1)

    if entries_count == 0:
        log("Script failed, no diagnostic logs found.")
        sys.exit(1)

    # ----------------------------------------------------------------
    # 6. Find entry with the greatest 'Id'
    # ----------------------------------------------------------------
    members = diag_data.get("Members", [])
    if not members:
        log("Script failed, 'Members' array is empty or missing.")
        sys.exit(1)

    id_greatest = -1
    entry_greatest_index = -1

    for i, entry in enumerate(members):
        try:
            entry_id = int(entry.get("Id", -1))
            if entry_id >= id_greatest:
                id_greatest = entry_id
                entry_greatest_index = i
        except ValueError:
            continue

    if entry_greatest_index == -1:
        log("Script failed, no valid numeric 'Id' found in the log entries.")
        sys.exit(1)

    if DEBUG:
        log(f"Greatest entry ID: {id_greatest}")

    # ----------------------------------------------------------------
    # 7. Verify the entry is of type "AllLogs"
    # ----------------------------------------------------------------
    entry_type = members[entry_greatest_index].get("OEMDiagnosticDataType", "")
    if entry_type != "AllLogs":
        log(f"Entry ID {id_greatest} is of type '{entry_type}' instead of 'AllLogs'.")
        sys.exit(1)

    # ----------------------------------------------------------------
    # 8. Download the attachment
    # ----------------------------------------------------------------
    log("Downloading AllLogs...")

    attachment_uri = members[entry_greatest_index].get("AdditionalDataURI", "")
    if not attachment_uri:
        log("No 'AdditionalDataURI' found for the chosen entry.")
        sys.exit(1)

    download_url = f"{PROTOCOL}://{bmc_ip}:{PORT}{attachment_uri}"
    filename = (
        f"{bmc_ip}_{datetime.now().strftime('%Y-%m-%dT%H-%M-%S')}_all_logs.tar.xz"
    )

    try:
        with http_request_with_retries(
            "get", download_url, auth=(bmc_username, bmc_password), stream=True
        ) as r:
            check_response_success(r, f"Failed to download logs at {download_url}")
            with open(filename, "wb") as f:
                for chunk in r.iter_content(chunk_size=8192):
                    if chunk:
                        f.write(chunk)
    except Exception as e:
        log(f"Exception while downloading logs: {e}")
        sys.exit(1)

    log(f"All logs downloaded as {filename}")


if __name__ == "__main__":
    main()
