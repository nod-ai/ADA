#!/bin/bash
#Copyright(C) 2024 Advanced Micro Devices, Inc. All rights reserved.
 
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
 
# Prompt for BMC Username if not set
if [ -z "$BMC_USERNAME" ]; then
    read -p "Enter BMC Username: " BMC_USERNAME
fi
 
# Prompt for BMC Password if not set
if [ -z "$BMC_PASSWORD" ]; then
    read -sp "Enter BMC Password: " BMC_PASSWORD
    echo
fi
 
# Prompt for BMC IP Address if not set and validate format
while [[ ! "$BMC_IP" =~ ^[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+$ ]]; do
    read -p "Enter valid BMC IP Address: " BMC_IP
done
 
# Step 1: Collect Diagnostic Data
log "Collecting diagnostic data..."
TASK_RESPONSE=$(curl -s -k -u "$BMC_USERNAME:$BMC_PASSWORD" \
    "https://$BMC_IP/redfish/v1/Oem/Supermicro/MI300X/Systems/UBB/LogServices/DiagLogs/Actions/LogService.CollectDiagnosticData" \
    -X POST -d '{"DiagnosticDataType":"OEM", "OEMDiagnosticDataType" : "AllLogs"}')
check_curl_error "Failed to collect diagnostic data."
 
# Extract Task ID
TASK_ID=$(echo "$TASK_RESPONSE" | grep -oP '(?<=Tasks/)[^"]*')
log "Task ID: $TASK_ID"
 
# Step 2: Poll for Task Completion with timeout
log "Checking task status..."
TASK_STATE=""
max_retries=40
retry_count=0
while [ "$TASK_STATE" != "Completed" ]; do
    if [ "$retry_count" -ge "$max_retries" ]; then
        log "Task did not complete within the expected timeframe."
        exit 1
    fi
    TASK_STATE=$(curl -s -k -u "$BMC_USERNAME:$BMC_PASSWORD" \
        -X GET "https://$BMC_IP/redfish/v1/Oem/Supermicro/MI300X/TaskService/Tasks/$TASK_ID" \
        | grep -oP '(?<=TaskState": ")[^"]*')
    log "Current State: $TASK_STATE"
    if [ "$TASK_STATE" != "Completed" ]; then
        log "Task in progress... waiting 30 seconds."
        sleep 30
        retry_count=$((retry_count + 1))
    fi
done
log "Task completed!"
 
# Extract Entry ID after completion
ENTRY_ID=$(curl -s -k -u "$BMC_USERNAME:$BMC_PASSWORD" \
    -X GET "https://$BMC_IP/redfish/v1/Oem/Supermicro/MI300X/TaskService/Tasks/$TASK_ID" \
    | grep -oP '(?<=Entries/)[^"]*')
#log "Entry ID: $ENTRY_ID"
 
# Step 3: Download All Logs to logs directory
OUTPUT_DIR="logs"
mkdir -p "$OUTPUT_DIR"
log "Downloading all logs..."
curl -s -k -u "$BMC_USERNAME:$BMC_PASSWORD" \
    "https://$BMC_IP/redfish/v1/Oem/Supermicro/MI300X/Systems/UBB/LogServices/DiagLogs/Entries/$ENTRY_ID/attachment" > "$OUTPUT_DIR/all_logs.tar.xz"
check_curl_error "Failed to download logs."
log "All logs downloaded as $OUTPUT_DIR/all_logs.tar.xz"
