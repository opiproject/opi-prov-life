# Notes from today meeting April 26

## Housekeeping
- Please use our repo https://github.com/opiproject/opi-prov-life
- Please use our slack https://opi-project.slack.com/archives/C0342L6T7EC
- meeting will be shortened to 45 minutes

## Docs
- please look at https://opentelemetry.io/docs/
- please look at https://github.com/silicom-ltd/uBMC
- please look at https://opencomputeproject.github.io/onie/overview/
- please look at https://fidoalliance.org/intro-to-fido-device-onboard/
- please look at https://www.rfc-editor.org/rfc/pdfrfc/rfc8572.txt.pdf (Secure Zero Touch Provisioning (SZTP))

## Monitoring & Telemetry
- We are still proceeding with adoption idea of https://opentelemetry.io/ for OPI
- Is there integration of OTEL with kvm or esx ?
- Example for OTEL collector with SPDK (json-rpc) and System monitoring on github
- Is it possible to produce the same example without DPU ? for example using KVM ?
- Use case of standalone DPU, not attached to server. Still runs OTEL collector
- For BMC, otel collector will use redfish input to get BMC specific metrics (like temp, power,...)
- We need to create a diagram for our examples on OTEL colector
- Tracing inside DPU/IPU (more tight SDK integration into our service and IPDK), streaming to zipkin/jaeger
- We are discussing OTEL collector deployment models:
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

## Initial DPU/IPU discovery and Provisioning
- Lee presented FDO high level mechanism and challanges (uploading to githuib...)
- Where the ZTP/FDO is running ? Suggestion not to run on BMC, but on the ARM cores
- Certificates and Authentication
  - What is assumed comes from factory ?
  - Can we fall into some defaults ?
  - Storage space limitations in DPU/IPU for CA
- Need to start working on use cases and deployments to move forward with ZTP

## Lifecycle & Updates
- nothing

## Scheduling
- nothing

## Action items:
- Dan: please drive SPDK tracing discussion with OTEL SDK
- Kyle & Dan - please find contact from MEV provisioning team
- Prasun - please see if Marvel can present something in provisioning space
- RedHat: please start participate in provisioning subgroup