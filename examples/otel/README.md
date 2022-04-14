# OpenTelemetry Collector + SPDK Proxy Example

## Prerequisites: 
- Docker >= 20.10.x

### Example environment components:

- [OpenTelemetry Collector](https://github.com/open-telemetry/opentelemetry-collector-contrib) - Collect, process, and export telemetry
- [Grafana](https://github.com/grafana/grafana) - Data visualization
- [InfluxDB](https://github.com/influxdata/influxdb) - Telemetry datastore

### Getting started:

1. Run `docker-compose up -d`
2. Open `http://localhost:3000/explore` for querying InfluxDB

### TODO

[ ] Setup SPDK and Proxy

[ ] Configure OTel Collector receiver for JSON events from SPDK Proxy

[ ] Create a simple example dashboard
