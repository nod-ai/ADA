#!/bin/bash
#Copyright(C) 2024 Advanced Micro Devices, Inc. All rights reserved.

DEBUG=false

# Function for logging with timestamps
log() {
    echo "$(date +'%Y-%m-%d %H:%M:%S') - $1"
}

# Function to check curl command success
check_curl_error() {
    if [ $? -ne 0 ]; then
        log "$1"
        exit 1
    fi
}

# Wait for all UBB SMC tasks to complete
tasksWait() {
    TIMEOUT_SECONDS=$((25 * 60)) # 25 minutes
    INTERVAL=3

    # Test to make sure there is only one task
    tasks=$(curl -k -s -u "${BMC_USERNAME}:${BMC_PASSWORD}" -X GET \
        -H 'Content-Type: application/json' -H 'Accept: application/json' \
        https://${BMC_IP}/redfish/v1/Oem/Supermicro/MI300X/TaskService/Tasks |
        python3 -c "import sys, json; print(json.load(sys.stdin)['Members@odata.count'])")

    if ${DEBUG}; then
        log "Tasks: ${tasks}"
    fi

    if [[ ${tasks} != "0" ]]; then
        for task in $(seq 0 $((tasks - 1))); do
            newlineNeeded=0
            elapsedTime=0

            # Poll the task status until it's completed or failed
            while true; do
                STATUS=$(curl -s -k -u "${BMC_USERNAME}:${BMC_PASSWORD}" \
                    -H 'Content-Type: application/json' -H 'Accept: application/json' \
                    https://${BMC_IP}/redfish/v1/Oem/Supermicro/MI300X/TaskService/Tasks/${task} |
                    python3 -c "import sys, json; print(json.load(sys.stdin)['TaskState'])")

                case "$STATUS" in
                "Completed")
                    if ${DEBUG}; then
                        printf "\n%s" "Task ${task} completed successfully."
                    fi

                    break
                    ;;
                "Failed")
                    log "Task ${task} failed."
                    exit 1
                    break
                    ;;
                "Running" | "New" | "Pending")
                    newlineNeeded=1
                    printf "\r%s" "Task ${task} still running, elapsed time ${elapsedTime}s"
                    sleep ${INTERVAL}
                    elapsedTime=$((elapsedTime + INTERVAL))
                    ;;
                *)
                    newlineNeeded=1
                    printf "\r%s" "Unknown task status: ${STATUS}, task ${task}, elapsed time ${elapsedTime}s"
                    sleep ${INTERVAL}
                    elapsedTime=$((elapsedTime + INTERVAL))
                    ;;
                esac

                if [ ${elapsedTime} -gt ${TIMEOUT_SECONDS} ]; then
                    log "Task ${task} fails to complete in $((TIMEOUT_SECONDS / 60)) minutes."
                    exit 1
                fi
            done

            if [[ ${newlineNeeded} != 0 ]]; then
                printf "\n"
            fi
        done
    fi
}

# Prompt for BMC Username if not set
if [ -z "${BMC_USERNAME}" ]; then
    read -p "Enter BMC Username: " BMC_USERNAME
fi

# Prompt for BMC Password if not set
if [ -z "${BMC_PASSWORD}" ]; then
    read -s -p "Enter BMC Password: " BMC_PASSWORD
    echo
fi

# Prompt for BMC IP Address if not set and validate format
while [[ ! ${BMC_IP} =~ ^[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+$ ]]; do
    read -p "Enter valid BMC IP Address: " BMC_IP
done

# Check if the username and password are correct
curl -s -u "${BMC_USERNAME}:${BMC_PASSWORD}" -k "https://${BMC_IP}/redfish/v1/Managers/1" |
    python3 -c '
import sys, json

data = json.load(sys.stdin)

if "error" in data:
    sys.exit(1)

sys.exit(0)'

if [ $? -ne 0 ]; then
    echo "Authentication failed, check IP address, username, and password"
    echo "IP address used: ${BMC_IP}"
    echo "Username used: ${BMC_USERNAME}"
    exit 1
fi

# Step 1: Collect Diagnostic Data
log "Collecting AllLogs file..."
taskResponse=$(curl -s -k -u "${BMC_USERNAME}:${BMC_PASSWORD}" \
    "https://${BMC_IP}/redfish/v1/Oem/Supermicro/MI300X/Systems/UBB/LogServices/DiagLogs/Actions/LogService.CollectDiagnosticData" \
    -X POST -d '{"DiagnosticDataType":"OEM", "OEMDiagnosticDataType" : "AllLogs"}')
check_curl_error "Failed to collect diagnostic data."

# Extract Task ID
TASKS=$(echo "${taskResponse}" | grep -oP '(?<=Tasks/)[^"]*')

if [[ ${TASKS} == "" ]]; then
    log "Script failed, task ID has no value"
    exit 1
fi

# Step 2: Poll for all tasks to complete
tasksWait

entries=$(curl -k -s -u "${BMC_USERNAME}:${BMC_PASSWORD}" -X GET \
    -H 'Content-Type: application/json' -H 'Accept: application/json' \
    https://${BMC_IP}/redfish/v1/Oem/Supermicro/MI300X/Systems/UBB/LogServices/DiagLogs/Entries |
    python3 -c "import sys, json; print(json.load(sys.stdin)['Members@odata.count'])")

if [[ ${entries} == "0" ]]; then
    log "Script failed, logs not found"
    exit 1
fi

idGreatest=0
entryGreatest=0

# Step 3: Go through the entries and find the one with the greatest ID
for ((entry = $((entries - 1)); entry >= 0; entry--)); do
    entryId=$(curl -k -s -u "${BMC_USERNAME}:${BMC_PASSWORD}" -X GET \
        -H 'Content-Type: application/json' -H 'Accept: application/json' \
        https://${BMC_IP}/redfish/v1/Oem/Supermicro/MI300X/Systems/UBB/LogServices/DiagLogs/Entries |
        python3 -c "import sys, json; print(json.load(sys.stdin)['Members'][${entry}]['Id'])")

    if [ ${idGreatest} -le ${entryId} ]; then
        idGreatest=${entryId}
        entryGreatest=${entry}
    fi
done

entryId=${entryGreatest}
if ${DEBUG}; then
    log "Entry ID: ${entryId}"
fi

# Step 4: Check to make sure the entry is an AllLogs
entryType=$(curl -k -s -u "${BMC_USERNAME}:${BMC_PASSWORD}" -X GET \
    -H 'Content-Type: application/json' -H 'Accept: application/json' \
    https://${BMC_IP}/redfish/v1/Oem/Supermicro/MI300X/Systems/UBB/LogServices/DiagLogs/Entries |
    python3 -c "import sys, json; print(json.load(sys.stdin)['Members'][${entryGreatest}]['OEMDiagnosticDataType'])")

if [[ ${entryType} != "AllLogs" ]]; then
    log "Entry ID ${entryId} is of type ${entryType}, rather than AllLogs"
    exit 1
fi

# Step 5: Download All Logs to logs directory
log "Downloading AllLogs..."

attachment=$(curl -k -s -u "${BMC_USERNAME}:${BMC_PASSWORD}" -X GET \
    -H 'Content-Type: application/json' -H 'Accept: application/json' \
    https://${BMC_IP}/redfish/v1/Oem/Supermicro/MI300X/Systems/UBB/LogServices/DiagLogs/Entries |
    python3 -c "import sys, json; print(json.load(sys.stdin)['Members'][${entryId}]['AdditionalDataURI'])")

filename="${BMC_IP}_$(date +"%Y-%m-%dT%H-%M-%S")_all_logs.tar.xz"
curl -k -s -u "${BMC_USERNAME}:${BMC_PASSWORD}" -X GET \
    https://${BMC_IP}${attachment} >${filename}
check_curl_error "Failed to download logs."
log "All logs downloaded as ${filename}"

