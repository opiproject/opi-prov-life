# Notes from today meeting May 3

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
- Raja from juniper.net raised concern they had regarding imaturity of OTEL when they evaluated this
  - no objection go back and re-evaluate this again if/when matures
- Granville didn't complete POC with OTEL yet, PR is still open 
- No answer yet on : Is there integration of OTEL with kvm or esx ?
- Tedd uploaded initial drawing on OTEL collector to github ../architecture/OPITelemetryArchitecture.drawio.png
- Tracing inside DPU/IPU (more tight SDK integration into our service and IPDK), streaming to zipkin/jaeger

## Initial DPU/IPU discovery and Provisioning
- Lee uploaded FDO presentation to github 
- Elad presented use-case subgroup focus - firewall, thi is not enough for provisioning
- Raja from juniper.net mentioned JESP they developed (look for blog)
- Need to start working on use cases and deployments to move forward with ZTP
  - mgmgt cable connected/not
  - far edge is difficult - not proprity
  - entherprise as well as SCPs are definitely use cases
  - with and without internet access (rendezvous server)
  - ...

## Lifecycle & Updates
- nothing

## Scheduling
- nothing

## Action items:
- Dan: please drive SPDK tracing discussion with OTEL SDK
- Kyle & Dan - please find contact from MEV provisioning team
- Prasun - please see if Marvel can present something in provisioning space
- RedHat: please start participate in provisioning subgroup
- Boris: start new page/slack to gather use cases for provisioning