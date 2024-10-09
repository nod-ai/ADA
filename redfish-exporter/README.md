# README #

The AMD Redfish Exporter is a Go server that subscribes to events and telemetry from a server with Redfish service enabled. It then exposes this data on a Prometheus-compatible metrics route for scraping. Additionally, it can initiate specific Slurm actions in response to certain events.

## Getting Started

### Prerequisites
- Docker
- Go 1.23+

### Running the Redfish Mock Server Locally
To test the exporter locally, you can deploy a Redfish mock server using Docker:
```bash
docker run -d --net="host" --name redfish-mock dmtf/redfish-mockup-server:latest --port=8000
```

> [!IMPORTANT]
> The `--net="host"` option allows the containerized application to share the host's networking, facilitating access to the listener endpoint. This configuration is essential if the applications aren't on the same network (e.g., when the listener app runs on the host).

#### Testing a Subscription
To manually create a subscription, create a JSON file (e.g. `data.json`) with the following content:
```json
{
    "Destination":"http://localhost:8080/",
    "Protocol":"Redfish",
    "EventFormatType":"Event",
    "IncludeOriginOfCondition":"true",
    "EventTypes":[
       "Alert",
       "ResourceRemoved",
       "ResourceAdded",
       "ResourceUpdated",
       "StatusChange"
    ]
 }
```

Send a POST request to register this subscription:
```bash
curl -d "@data.json" -X POST http://localhost:8000/redfish/v1/EventService/Subscriptions
```

Verify the subscription is successfully created:
```bash
curl http://localhost:8000/redfish/v1/EventService/Subscriptions
curl http://localhost:8000/redfish/v1/EventService/Subscriptions/[ID]
```

You can trigger a test event by creating a file (e.g., `testevent.json`) with the following content:
```json
{
    "EventType": "Alert",
    "EventId": "TestEvent",
    "EventTimestamp": "2024-09-30T10:10:10Z",
    "Severity": "Warning",
    "Message": "This is a test event message",
    "MessageId": "Basic.1.5.SomethingIsHappening",
    "MessageArgs": ["1", "slot 3"],
    "OriginOfCondition": "/redfish/v1/Systems/1/Storage"
}
```

Submit the test event with:
```bash
curl -d "@testevent.json" -X POST http://localhost:8000/redfish/v1/EventService/Actions/EventService.SubmitTestEvent
```

> [!WARNING]
> Unlike what's shown [here](https://github.com/DMTF/Redfish-Mockup-Server/blob/main/public-rackmount1/EventService/SubmitTestEventActionInfo/index.json) all the fields are required by the mockup server. otherwise, the request will result in a BadRequest 400. check [here](https://github.com/DMTF/Redfish-Mockup-Server/blob/2d39eb14122337ceab0712a9610b1cd37c65f487/redfishMockupServer.py#L169) for more info.

To delete a subscription, use its ID in a DELETE request:
```bash
curl -X DELETE http://localhost:8000/redfish/v1/EventService/Subscriptions/[ID]
```

### Running the app

Follow the [Slurm API setup guide](./api/README.md) to configure Slurm API integration.

Update the `.env` file with your Redfish server details.

Start the AMD Redfish Exporter:
```bash
go run .
```

The server listens on:
- Port 8080 for Redfish event handling.
- Port 2112 for Prometheus metrics.

> [!WARNING]
> Ensure the `Destination` in the `SUBSCRIPTION_PAYLOAD`  is set to `http://localhost:8080` for local testing.

The server will automatically subscribe to the Redfish servers defined in `REDFISH_SERVERS`.
Verify the subscription:
```bash
curl http://localhost:8000/redfish/v1/EventService/Subscriptions
```

Verify the event handling:
```bash
curl -d "@testevent.json" -X POST http://localhost:8000/redfish/v1/EventService/Actions/EventService.SubmitTestEvent
```

This will trigger a test event on the Redfish server, and you should see it reflected in the listener logs.

Currently, the exporter dynamically generates two metrics per Redfish event. To define additional metrics, modify the [metrics](./metrics/metrics.go) file.

To view the metrics:
```bash
curl "http://127.0.0.1:2112/metrics"
```

You should see output similar to:
```
# HELP RedFishEvents_processing_time Time taken to process events
# TYPE RedFishEvents_processing_time gauge
RedFishEvents_processing_time{EventType="Alert",SourceIP="127.0.0.1"} 1.728467499e+09
# HELP RedFishEvents_recieved Total number of events processed
# TYPE RedFishEvents_recieved counter
RedFishEvents_recieved{EventType="Alert",SourceIP="127.0.0.1"} 1
```
These metrics will update as more data is processed.

### Running with SSL

Generate a Self-Signed Certificate:
```bash
# Create a Private Key
openssl genrsa -out cert.key 4096

# Generate a Certificate Signing Request (CSR)
openssl req -new -key cert.key -out csr.pem

# Sign the CSR to create the SSL Certificate
openssl x509 -req -days 365 -in csr.pem -signkey cert.key -out cert.crt
```

Set the path to these certificates in your `.env` file and enable SSL by setting `USE_SSL` to `True`.

### Running the Mock Server in kubernetes

Deploy the Redfish Mockup Server in Kubernetes using the provided kustomize template:
```bash
kubectl apply -k dev/redfish-mock
```

Once deployed, expose the service locally:
```bash
kubectl port-forward -n silo services/redfish-mock 8000
```

You can then interact with the mock server:
```bash
curl localhost:8000/redfish/v1
```

Note that SSL can be enabled using the provided dev certificate and key and `--ssl` flag.

### Slurm Integration

1. Follow [this](./api/README.md) guide for generating the Slurm REST API.

2. Required configurations:

- **SLURM_TOKEN**: Obtain the token by running the following command on the Slurm control node (where slurmrestd is running):
    ```sh
    scontrol token username=root lifespan=18000
    ```
    (`lifespan` is specified in seconds)

- **SLURM_CONTROL_NODE**: The hostname or IP address of the Slurm control node where the REST server is running.

3. Run the exporter with `--enable-slurm` flag to enable slurm
    ```sh
    ./amd-redfish-exporter --enable-slurm
    ```

### Integration Testing

To run integration tests:
```bash
make integration-test
```

This will:
- Build the AMD Redfish Exporter
- Spin up two mock Redfish servers using Docker
- Test that events are correctly exported as Prometheus metrics.
- Clean up the test environment.
