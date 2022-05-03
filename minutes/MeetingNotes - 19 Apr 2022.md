# Notes from today meeting April 19

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
- We having active OTEL discussions in SLACK, please join
- Tomorrow on the core call, https://github.com/gramidt will join to discuss OTEL and answer questions
- We have 2 levels of OTEL we are discussing:
  - System monitoring (cpu,mem,nic,...) see https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/hostmetricsreceiver
  - Tracing inside DPU/IPU (more tight SDK integration into our service and IPDK), streaming to zipkin/jaeger
  - Am I missing another levels ?
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
- Andy presented Dell Servers provisioning mechanism (uploading to githuib...)
- Certificates and Authentication
  - What is assumed comes from factory ?
  - Can we fall into some defaults ?
  - Storage space limitations in DPU/IPU for CA
- What exactly can we standardize?
  - Is there a single method that fits all use cases? 
  - Should we enable different for different customers and use cases, depending on their existing provisioning method?

## Lifecycle & Updates
- nothing

## Scheduling
- nothing

## Action items:
- Kyle: please invite Dan.D to discuss how can we integrate OTEL and IPDK doscussion
- RedHat: please start participate in provisioning subgroup
- Anh Thu: please ask around and present how Marvel Octeon does provisioning today
- Andy: please look into FDO as well in addition to sZTP if it can complement sZTP
