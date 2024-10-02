# README #

The AMD Redfish Exporter is a Go server that subscribes to events and telemetry from a server with Redfish service enabled. It then exposes this data on a Prometheus-compatible metrics route for scraping.

### Getting Started ###

Ensure you have Go 1.23 installed.

To set up the server, run:
```bash
go run .
```

This will start a server that listens on port 8080 for events and port 2112 for metrics.

To test it out, you can use:
```bash
curl -X POST http://127.0.0.1:8080 \
     -H "Content-Type: application/json" \
     -H "Custom-Header: CustomValue" \
     -d '{"EventType":"Alert","EventId":"TestEvent","EventTimestamp":"2024-09-30T10:10:10Z","Severity":"Warning","Message":"This is a test event message","MessageId":"Basic.1.5.SomethingIsHappening","MessageArgs":["1","slot 3"],"OriginOfCondition":"/redfish/v1/Systems/1/Storage"}'
```
This will post some sample data.

Next, retrieve the metrics with:
```bash
curl "http://127.0.0.1:2112/metrics"
```
scroll to the end to see the metrics:

```
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 1
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
```
These metrics will update as you post more data.

### Running the Mock Server locally ###
To run the Redfish mock server locally, use the following `docker run` command:
```bash
docker run -d --net="host" --name redfish-mock dmtf/redfish-mockup-server:latest --port=8000
```

[!IMPORTANT]
The `--net="host"` option allows the program inside the Docker container to behave as if it's running directly on the host machine, enabling access to the listener endpoint. This is only needed if the applications are not on the same network (e.g., The listener app is running on the host).

You can create a subscription manually by creating a JSON file (e.g., `data.json`) with the following content:
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

Then send a POST request to create the subscription:
```bash
curl -d "@data.json" -X POST http://localhost:8000/redfish/v1/EventService/Subscriptions
```

Check if the subscription is created:
```bash
curl http://localhost:8000/redfish/v1/EventService/Subscriptions
curl http://localhost:8000/redfish/v1/EventService/Subscriptions/[ID]
```

To delete the subscription, send a DELETE request by specifying the subscription ID:
```bash
curl -X DELETE http://localhost:8000/redfish/v1/EventService/Subscriptions/[ID]
```

To trigger test events, create a JSON file (e.g., `testevent.json`) with the following content:
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

[!WARNING]
Unlike what's shown [here](https://github.com/DMTF/Redfish-Mockup-Server/blob/main/public-rackmount1/EventService/SubmitTestEventActionInfo/index.json) all the fields are required by the mockup server. otherwise, the request will result in a BadRequest 400. check [here](https://github.com/DMTF/Redfish-Mockup-Server/blob/2d39eb14122337ceab0712a9610b1cd37c65f487/redfishMockupServer.py#L169) for more info.

Then, use the following command to send the test event:
```bash
curl -d "@testevent.json" -X POST http://localhost:8000/redfish/v1/EventService/Actions/EventService.SubmitTestEvent
```

### Running the Mock Server in kubernetes ###
The Redfish Mockup Server handles Redfish requests against a mock setup. You can deploy it using the redfish-mock kustomize template, following the steps:
```bash
kubectl apply -k dev/redfish-mock
```

Once the server is running, expose the service locally using:
```bash
kubectl port-forward -n silo services/redfish-mock 8000
```
You can then query the Redfish endpoint with:
```bash
curl localhost:8000/redfish/v1
```
Note that SSL can be enabled using the provided dev certificate and key.

### Integration Testing

To run the integration tests, use: `make integration-test`

This will:
- Build the AMD Redfish Exporter
- Spin up two mock Redfish servers using Docker
- Test the exporter against these servers to ensure Redfish events are properly exported
- Clean up the environment afterward

