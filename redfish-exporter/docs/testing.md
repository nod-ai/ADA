# Mock Server and Integration Testing Documentation

## Running the Mock Redfish Server

The mock Redfish server simulates a Redfish-enabled device, allowing you to test the AMD Redfish Exporter without connecting to real hardware.

### Using Docker

1. Run the following command to start the mock Redfish server:

```bash
docker run -d --net="host" --name redfish-mock dmtf/redfish-mockup-server:latest --port=8000
```

Note: The `--net="host"` option allows the container to access the host network directly. This is useful for local testing but may not be suitable for all environments.

1. Verify that the server is running:

```bash
curl http://localhost:8000/redfish/v1/
```

You should receive a JSON response representing the root of the Redfish service.

### Using Kubernetes

1. Apply the Kubernetes configuration:

```bash
kubectl apply -k dev/redfish-mock
```

1. Forward the service port to your local machine:

```bash
kubectl port-forward -n silo services/redfish-mock 8000:8000
```

1. Verify that the server is running:

```bash
curl http://localhost:8000/redfish/v1/
```

## Integration Testing

Integration tests ensure that the AMD Redfish Exporter works correctly with the mock Redfish server and other components.

### Prerequisites

- Go 1.23 or later
- Docker
- Make

### Running the Integration Tests

1. Navigate to the project root directory.

2. Run the integration tests using the following command:

```bash
make integration-test
```

This command will:

- Build the AMD Redfish Exporter
- Start two mock Redfish servers using Docker
- Run the integration tests
- Clean up the test environment

1. Monitor the output for test results and any error messages.

### Test Coverage

The integration tests cover the following scenarios:

1. Redfish event subscription
2. Event reception and processing
3. Metric exposition
4. SLURM integration (if enabled)

[Placeholder: Add more specific test scenarios as they are implemented]

## Troubleshooting

If you encounter issues while running the mock server or integration tests, try the following:

- Ensure all prerequisites are installed and up to date.
- Check that no other services are using the required ports (8000 for the mock server, 8443 for the exporter).
- Review the Docker and Kubernetes logs for any error messages:

```bash
# For Docker
docker logs redfish-mock

# For Kubernetes
kubectl logs -n silo deployment/redfish-mock
```

- Verify that your firewall or security settings are not blocking the required network connections.

[Placeholder: Add more troubleshooting steps based on common issues encountered during testing]

If problems persist, please file an issue on the project's GitHub repository with detailed information about the error and your environment.
