# Monitoring & Telemetry

- OPI adoped https://opentelemetry.io/ for DPUs
- OPI goal is to pick 1 standard protocol that
  - all vendors can implement (both linux and non-linux based)
  - DPU consumers can integrate once in their existing monitoring systems and tools

- OpenTemetry suports those data sources
  - Traces
  - Metrics
  - Logs

## What is mandated by OPI

- OpenTemetry is made up of several main components:
  - Specification https://github.com/open-telemetry/opentelemetry-specification
  - Collector https://github.com/open-telemetry/opentelemetry-collector
  - SDKs (different programming languages), for example https://github.com/open-telemetry/opentelemetry-java)

- OPI is only mandating OTEL `Specification`
- SDK and Collector specific implementation are left to the users
  - They can be also from OTEL or other sources.

## Collector deploy options

![](/architecture/OPITelemetryArchitecture.drawio.png)

- OpenTemetry collector supports several deployments:
  - Deploy as side car inside every pod
  - Deploy another one as aggregator per Node
  - Deploy another one as super-aggregator per Cluster

- The benefits of having multiple collectors at different levels are:
  - Increased redundancy
  - Enrichment
  - Filtering
  - Separating trust domains
  - Batching
  - Sanitization

- Recommendation (reference)
  - micro-aggregator inside each DPU/IPU
  - macro-aggregator between DPUs
    - macro-aggregator can run on the host with DPU/IPU or on a separate host

## System Monitoring

- System monitoring (cpu,mem,nic,...) 
  - see as example https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/hostmetricsreceiver

- BMC monitoring (temp, power)
  - Redfish extention for OTEL collector can be used to collect HW/BMC related metrics like temperature, power and others...

- TBD
  - OPI wants to define which telemetries are mandatory for each vendor to implement and which are optional

## Tracing

- Tracing inside DPU/IPU (more tight SDK integration into our service and IPDK), streaming to zipkin/jaeger
- TODO: need more details and examples

## Logging

- For example https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/syslogreceiver
- TODO: need more details and examples

## Examples

- see https://github.com/opiproject/opi-prov-life/tree/main/examples
- Example for OTEL collector with SPDK (json-rpc) and System monitoring on github
- Same example without HW DPU, for example using KVM
- Example using eBPF
  - https://www.splunk.com/en_us/blog/devops/announcing-the-donation-of-the-opentelemetry-ebpf-collector.html
- TODO: Example with IPDK as well

## questions to  (eventually remove this section)

- Is there integration of OTEL with kvm or esx ?
- Use case of standalone DPU, not attached to server. Still runs OTEL collector
