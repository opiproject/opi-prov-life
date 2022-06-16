# Notes from today meeting April 12

## Housekeeping
- transition to new repo completed https://github.com/opiproject/opi-prov-life
- meeting will be renamed from Diamond Bluff to OPI
- meeting will be shortened to 45 minutes

## Docs
- please look at https://opentelemetry.io/docs/
- please look at https://github.com/silicom-ltd/uBMC
- please look at https://opencomputeproject.github.io/onie/overview/
- please look at https://fidoalliance.org/intro-to-fido-device-onboard/
- please look at https://www.rfc-editor.org/rfc/pdfrfc/rfc8572.txt.pdf (Secure Zero Touch Provisioning (SZTP))

## Monitoring & Telemetry
- We are still proceeding with adoption idea of https://opentelemetry.io/ for OPI
- Next week one of the OTEL maintainers from F5 will come to discuss OTEL with us

## Initial DPU/IPU discovery and Provisioning
- Eyal + Elad presented how Silicom did provisioning with uBMC
- Gene presented compare table betwwen PXE / sZTP / FDO
- Tedd raised a good point whether Host BMC and DPU BMC have to talk or not
  - some feedback was received that they do
- Anh Thu raised a good point if we can come up with single provisioning method
  - Maybe for different use cases we will need different methods?

## Lifecycle & Updates
- nothing

## Scheduling
- nothing

## Action items:
- Kyle: please start IPDK doscussion on OpenTelemetry (same action from last week)
- Eyal: please think how would you do discovery and provisioning for uBMC differently today (2-nd gen uBMC)
- Gene: please upload compare table to github
- Elad+Eyal: please upload presentation to github
- Andy+Lee: please present how provisioning works today in Dell
- Tedd: please start working on security models, Gene and Anh Thu - please join
