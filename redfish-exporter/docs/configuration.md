# Configuration Guide

## Environment Variables

The AMD Redfish Exporter can be configured using the following environment variables:

- `UPDATED`: Last update date of the exporter
- `DESCRIPTION`: Description of the exporter
- `LISTENER_IP`: IP address to listen on (default: "127.0.0.1")
- `LISTENER_PORT`: Port to listen on for Redfish events (default: "8080")  (API events?)
- `USE_SSL`: Enable SSL/TLS (default: "false")
- `CERTFILE`: Path to SSL certificate file
- `KEYFILE`: Path to SSL key file
- `METRICS_PORT`: Port to expose Prometheus metrics (default: "2112")
- `SLURM_TOKEN`: SLURM authentication token
- `SLURM_CONTROL_NODE`: SLURM control node address
- `SUBSCRIPTION_PAYLOAD`: JSON string containing Redfish subscription details
- `REDFISH_SERVERS`: JSON array of Redfish server configurations
- `TRIGGER_EVENTS`: JSON array of events that trigger actions

## Redfish Server Configuration

The `REDFISH_SERVERS` environment variable should contain a JSON array of objects with the following structure:

```json
[
  {
    "ip": "https://redfish-server1.example.com",
    "username": "admin",
    "password": "password123",
    "loginType": "basic",
    "slurmNode": "compute-node-1"
  },
  {
    "ip": "https://redfish-server2.example.com",
    "username": "admin",
    "password": "password456",
    "loginType": "basic",
    "slurmNode": "compute-node-2"
  }
]
```

**NOTE:** Ensure the node name provided in `REDFISH_SERVERS` matches the SLURM node configuration. If the SLURM node name is entered incorrectly the node drain operation will fail.

## Subscription Payload

The `SUBSCRIPTION_PAYLOAD` environment variable should contain a JSON object with the following structure:

```json
{
  "Destination": "https://exporter.example.com:8080",
  "EventTypes": ["Alert", "ResourceAdded", "ResourceRemoved", "ResourceUpdated", "StatusChange"],
  "Context": "ExporterSubscription",
  "Protocol": "Redfish"
}
```

## Trigger Events

The `TRIGGER_EVENTS` environment variable should contain a JSON array of objects with the following structure:

```json
[
  {
    "Severity": "Critical",
    "Action": "DrainNode"
  },
  {
    "Severity": "Warning",
    "Action": "LogEvent"
  }
]
```

## Formatting JSON Objects

For environment variables that require JSON objects (such as `REDFISH_SERVERS`, `SUBSCRIPTION_PAYLOAD`, and `TRIGGER_EVENTS`), you can use the following one-liner to properly fomat a JSON object for use in `.env`:

```bash
echo MY_JSON=\'$(cat "path/to/your/json_file.json" | tr -d '\n' | sed "s/'/'\\\\''/g")\'  >> .env
```

Replace `MY_JSON` with the appropriate environment variable name, and `path/to/your/json_file.json` with the path to your JSON file.

### Example

- Create a JSON file with the required configuration. For example, `redfish_servers.json`:

```json
[
 {
   "ip": "https://redfish-server1.example.com",
   "username": "admin",
   "password": "password123",
   "loginType": "basic",
   "slurmNode": "compute-node-1"
 },
 {
   "ip": "https://redfish-server2.example.com",
   "username": "admin",
   "password": "password456",
   "loginType": "basic",
   "slurmNode": "compute-node-2"
 }
]
```

- Use the following one-liner to add the properly formatted environment variable to the `.env` configuration file:

```bash
echo REDFISH_SERVERS=\'$(cat "redfish_servers.json" | tr -d '\n' | sed "s/'/'\\\\''/g")\' >> .env
```

This command does the following:

- `cat "redfish_servers.json"`: Reads the content of the JSON file
- `tr -d '\n'`: Removes all newline characters, making it a single line
- `sed "s/'/'\\\\''/g"`: Escapes any single quotes in the JSON
- `echo <command> >> .env` Appends the output to the .env file

Repeat this process for other JSON object variables like `SUBSCRIPTION_PAYLOAD` and `TRIGGER_EVENTS`.

## SSL/TLS Configuration

To enable SSL/TLS:

1. Set `USE_SSL` to "true"
2. Provide paths to your SSL certificate and key files using `CERTFILE` and `KEYFILE` respectively

## SLURM Integration

To enable SLURM integration:

1. Provide the SLURM authentication token in `SLURM_TOKEN`
2. Specify the SLURM control node address in `SLURM_CONTROL_NODE`
