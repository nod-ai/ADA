# Metrics Documentation

The AMD Redfish Exporter exposes the following Prometheus-compatible metrics:

## Event Metrics

### `RedFishEvents_recieved`

- **Type**: Counter
- **Description**: Total number of Redfish events received
- **Labels**:
  - `SourceIP`: IP address of the Redfish server that sent the event
  - `Severity`: Severity of the Redfish event

### `RedFishEvents_processing_time`

- **Type**: Gauge
- **Description**: Time taken to process Redfish events in seconds
- **Labels**:
  - `SourceIP`: IP address of the Redfish server that sent the event
  - `EventType`: Type of the Redfish event

## Accessing Metrics

Metrics are exposed on the `/metrics` endpoint. By default, this is available at `http://localhost:2112/metrics`.

## Prometheus Configuration

To scrape metrics from the AMD Redfish Exporter, add the following job to your Prometheus configuration:

```yaml
scrape_configs:
  - job_name: 'amd_redfish_exporter'
    static_configs:
      - targets: ['localhost:2112']
```

[Placeholder: Additional Prometheus config options?]
