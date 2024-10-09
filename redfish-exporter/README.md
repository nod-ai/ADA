# AMD Redfish Exporter

The AMD Redfish Exporter is a Go server that subscribes to events and telemetry from servers with Redfish service enabled. It exposes this data on a Prometheus-compatible metrics route for scraping. The exporter also integrates with SLURM to drain nodes when critical hardware events are detected.

## Features

- Subscribes to Redfish events from multiple servers
- Processes and filters events based on configurable criteria
- Exposes Prometheus-compatible metrics
- Integrates with SLURM for node management
- Supports SSL/TLS for secure communication

## Installation

### Prerequisites

- Go 1.23 or later
- SLURM (for node management functionality)

### Building from source

- Clone the repository:

```bash
git clone https://github.com/nod-ai/ADA.git
cd ADA/redfish-exporter
```

- Build the project:

```bash
make build
```

- The binary will be created as `amd-redfish-exporter` in the current directory.

### Using Docker

A Dockerfile is provided to build and run the AMD Redfish Exporter in a container:

- Build the Docker image:

```bash
docker build -t amd-redfish-exporter .
```

- Run the container:

```bash
docker run -p 8080:8080 -p 2112:2112 amd-redfish-exporter
```

## Configuration

The AMD Redfish Exporter is configured using environment variables or a configuration file. Here are the main configuration options:

- `LISTENER_IP`: IP address to listen on (default: "127.0.0.1")
- `LISTENER_PORT`: Port to listen on for Redfish events (default: "8080")
- `METRICS_PORT`: Port to expose Prometheus metrics (default: "2112")
- `USE_SSL`: Enable SSL/TLS (default: "false")
- `CERTFILE`: Path to SSL certificate file (required if USE_SSL is true)
- `KEYFILE`: Path to SSL key file (required if USE_SSL is true)
- `SLURM_TOKEN`: SLURM authentication token
- `SLURM_CONTROL_NODE`: SLURM control node address
- `REDFISH_SERVERS`: JSON array of Redfish server configurations
- `SUBSCRIPTION_PAYLOAD`: JSON object defining the Redfish subscription parameters

For a complete list of configuration options, refer to the [Configuration Guide](docs/configuration.md).

## Usage

1. Set up the configuration as described in the [Configuration](#configuration) section.

2. Run the AMD Redfish Exporter:

```bash
./amd-redfish-exporter
```

1. The exporter will start subscribing to events from the configured Redfish servers and expose metrics on the specified metrics port.

2. Configure your Prometheus server to scrape this endpoint: `http://[server IP]:2112/metrics`.

For more detailed usage instructions, see the [User Guide](docs/user-guide.md).

## Metrics

The exporter provides the following metrics:

- `RedFishEvents_received`: Counter of Redfish events received
- `RedFishEvents_processing_time`: Gauge of event processing time

For a complete list of available metrics and their descriptions, see the [Metrics Documentation](docs/metrics.md).

## Testing

See [testing.md](docs/testing.md) for information on using the Mock Server and Integration Testing.

## Support

If you encounter any issues or have questions, please file an issue.
