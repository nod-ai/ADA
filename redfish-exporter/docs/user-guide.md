# User Guide

## Installation

### Prerequisites

- Go 1.23 or later
- SLURM (for node management functionality)

### Building from Source

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

- Build the Docker image:

```bash
docker build -t amd-redfish-exporter .
```

- Run the container:

```bash
docker run -p 8080:8080 -p 2112:2112 amd-redfish-exporter
```

## Configuration

Refer to the [Configuration Guide](configuration.md) for detailed information on how to configure the AMD Redfish Exporter.

## Running the Exporter

- Set up the configuration as described in the Configuration Guide.
- Run the AMD Redfish Exporter:

 ```bash
 ./amd-redfish-exporter
 ```

1. The exporter will start subscribing to events from the configured Redfish servers and expose metrics on the specified metrics port.

## Viewing Metrics

To view the metrics, navigate to `http://localhost:2112/metrics` in your web browser or use a Prometheus server to scrape this endpoint.

## Troubleshooting

[Placeholder]

## Updating

[Placeholder]

## Best Practices

[Placeholder]
