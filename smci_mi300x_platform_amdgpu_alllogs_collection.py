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

# Suppress InsecureRequestWarning if using verify=False in requests
import urllib3
urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)

# --------------------------------------------------------------------
# Variables for porting between platforms
# --------------------------------------------------------------------
REDFISH_OEM = "redfish/v1/Oem/Supermicro/MI300X"
REDFISH_MANAGER_BMC = "redfish/v1/Managers/1"

# --------------------------------------------------------------------
# Setup argument parsing for optional debug mode
# --------------------------------------------------------------------
parser = argparse.ArgumentParser(description="Collect AllLogs from BMC via Redfish")
parser.add_argument('--debug', action='store_true', help='Enable debug output')
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
    ip_pattern = re.compile(r'^\d+\.\d+\.\d+\.\d+$')

    while not ip_pattern.match(bmc_ip):
        bmc_ip = input("Enter valid BMC IP Address: ")

    return bmc_ip

def authenticate(bmc_ip, bmc_username, bmc_password):
    """
    Check if the provided IP, username, and password are valid by querying
    the Redfish Managers resource.
    """
    url = f"https://{bmc_ip}/{REDFISH_MANAGER_BMC}"
    try:
        response = requests.get(url, verify=False, auth=(bmc_username, bmc_password))
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
        response = requests.get(url, verify=False, auth=(bmc_username, bmc_password))
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
                response = requests.get(task_url, verify=False, auth=(bmc_username, bmc_password))
                check_response_success(response, f"Failed to query status of task {task_index}.")
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
                print(f"\rTask {task_index} still running, elapsed time {elapsed_time}s", end='')
                sys.stdout.flush()
                time.sleep(INTERVAL)
                elapsed_time += INTERVAL
            else:
                newline_needed = True
                print(f"\rUnknown task status: {status}, task {task_index}, elapsed time {elapsed_time}s", end='')
                sys.stdout.flush()
                time.sleep(INTERVAL)
                elapsed_time += INTERVAL

            if elapsed_time > TIMEOUT_SECONDS:
                log(f"Task {task_index} failed to complete in {TIMEOUT_SECONDS // 60} minutes.")
                sys.exit(1)

        if newline_needed:
            print()  # Move to the next line


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
    # 3. Collect diagnostic data (AllLogs)
    # ----------------------------------------------------------------
    log("Collecting AllLogs file...")
    collect_url = (f"https://{bmc_ip}/{REDFISH_OEM}/Systems/UBB/"
                   "LogServices/DiagLogs/Actions/LogService.CollectDiagnosticData")
    payload = {
        "DiagnosticDataType": "OEM",
        "OEMDiagnosticDataType": "AllLogs"
    }

    try:
        response = requests.post(collect_url, verify=False, auth=(bmc_username, bmc_password),
                                 json=payload)
        check_response_success(response, "Failed to collect diagnostic data.")
        task_response_text = response.text
    except Exception as e:
        log(f"Exception while collecting diagnostic data: {e}")
        sys.exit(1)

    # Extract Task ID from the response text
    match = re.search(r'Tasks/([^"]+)', task_response_text)
    if not match:
        log("Script failed, task ID has no value")
        sys.exit(1)
    tasks_id = match.group(1)  # Though not directly used, we confirm it exists.

    # ----------------------------------------------------------------
    # 4. Wait for tasks to complete
    # ----------------------------------------------------------------
    tasks_wait(bmc_ip, bmc_username, bmc_password)

    # ----------------------------------------------------------------
    # 5. Get count of diagnostic log entries
    # ----------------------------------------------------------------
    entries_url = (f"https://{bmc_ip}/{REDFISH_OEM}/Systems/UBB/"
                   "LogServices/DiagLogs/Entries")
    try:
        response = requests.get(entries_url, verify=False, auth=(bmc_username, bmc_password))
        check_response_success(response, "Failed to get diagnostic log entries.")
        diag_data = response.json()
        entries_count = diag_data.get("Members@odata.count", 0)
    except Exception as e:
        log(f"Exception while retrieving log entries: {e}")
        sys.exit(1)

    if entries_count == 0:
        log("Script failed, logs not found")
        sys.exit(1)

    # ----------------------------------------------------------------
    # 6. Find entry with the greatest 'Id'
    # ----------------------------------------------------------------
    # We already have the JSON data in 'diag_data', which has a 'Members' list
    members = diag_data.get("Members", [])
    if not members:
        log("Script failed, no Members in diag data.")
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
        log("Script failed, could not find a valid entry ID.")
        sys.exit(1)

    if DEBUG:
        log(f"Greatest entry ID: {id_greatest}")

    # ----------------------------------------------------------------
    # 7. Verify the entry is of type "AllLogs"
    # ----------------------------------------------------------------
    entry_type = members[entry_greatest_index].get("OEMDiagnosticDataType", "")
    if entry_type != "AllLogs":
        log(f"Entry ID {id_greatest} is of type {entry_type}, rather than AllLogs")
        sys.exit(1)

    # ----------------------------------------------------------------
    # 8. Download the attachment
    # ----------------------------------------------------------------
    log("Downloading AllLogs...")

    attachment_uri = members[entry_greatest_index].get("AdditionalDataURI", "")
    if not attachment_uri:
        log("No AdditionalDataURI found for the chosen entry.")
        sys.exit(1)

    download_url = f"https://{bmc_ip}{attachment_uri}"
    filename = f"{bmc_ip}_{datetime.now().strftime('%Y-%m-%dT%H-%M-%S')}_all_logs.tar.xz"

    try:
        with requests.get(download_url, verify=False, auth=(bmc_username, bmc_password), stream=True) as r:
            check_response_success(r, f"Failed to download logs at {download_url}.")
            with open(filename, 'wb') as f:
                for chunk in r.iter_content(chunk_size=8192):
                    if chunk:
                        f.write(chunk)
    except Exception as e:
        log(f"Exception while downloading logs: {e}")
        sys.exit(1)

    log(f"All logs downloaded as {filename}")


if __name__ == "__main__":
    main()

