# Host Metrics Receiver

The Host Metrics receiver generates metrics about the host system scraped
from various sources. This is intended to be used when the collector is
deployed as an agent.

Supported pipeline types: metrics

## Docs

* Plugins used
  * <https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/hostmetricsreceiver>

* Dockers used
  * <https://hub.docker.com/r/otel/opentelemetry-collector-contrib>
  * <https://hub.docker.com/_/influxdb>
  * <https://hub.docker.com/r/grafana/grafana>

## Getting started

The collection interval and the categories of metrics to be scraped can be
configured:

```yaml
hostmetrics:
  collection_interval: <duration> # default = 1m
  scrapers:
    <scraper1>:
    <scraper2>:
    ...
```

The available scrapers are:

| Scraper    | Supported OSs                | Description                                            |
|------------|------------------------------|--------------------------------------------------------|
| cpu        | All except Mac               | CPU utilization metrics                                |
| disk       | All except Mac               | Disk I/O metrics                                       |
| load       | All                          | CPU load metrics                                       |
| filesystem | All                          | File System utilization metrics                        |
| memory     | All                          | Memory utilization metrics                             |
| network    | All                          | Network interface I/O metrics & TCP connection metrics |
| paging     | All                          | Paging/Swap space utilization and I/O metrics
| processes  | Linux                        | Process count metrics                                  |
| process    | Linux & Windows              | Per process CPU, Memory, and Disk I/O metrics          |
