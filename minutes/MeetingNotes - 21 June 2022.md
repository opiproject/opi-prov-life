# Notes from today meeting June 21

## Housekeeping
- Please use our repo https://github.com/opiproject/opi-prov-life
- Please use our slack https://opi-project.slack.com/archives/C0342L6T7EC
- meeting will be shortened to 45 minutes

## Monitoring & Telemetry
- we adoped https://opentelemetry.io/ for OPI
- no more discussions at this point, unless specifically raised by participants

## Initial DPU/IPU discovery and Provisioning
- We continue to discuss common Inventory query https://github.com/opiproject/opi-prov-life/blob/main/INVENTORY.md
- We continue to discuss Discovery & Provisioning https://github.com/opiproject/opi-prov-life/blob/main/PROVISIONING.md
- We continue to discuss use cases https://github.com/opiproject/opi-prov-life/blob/main/USECASES.md
- We are heavily discussing in Slack, please join

- Juniper presented DEMO based on sZTP both client/agent and (bootsrap) server side. Recording is on slack

- Discussion where sZTP client/agent is running and how he gets there presented 3 options:
  - pre-configired from the factory (golden partition or not)
    - challanges for factories
  - ztp agent inside UEFI
    - challanges with Vendor specific supprot and writing code in UEFI, lack of flexability
    - have to solve security issues, how to set/change keys/sertificates in UEFI ? standard for that ?
    - Mutual authentication should be required for multi-tenant environments
  - chain load
    - network boot (can be HTTPs) min OS with sZTP agent inside to ramfs
    - run sztp agent
    - challanges who mantains min OS ? need Vendor drivers and support there
    - more flexble solution
    - still requires security solution
    - Mutual authentication should be required for multi-tenant environments

## Lifecycle & Updates
- nothing

## Scheduling
- nothing

## Action items:
- Juniper: present demo for SZTP next week 6/21
- Marvel: update manual provisioning method on MD
