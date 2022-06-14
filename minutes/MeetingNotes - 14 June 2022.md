# Notes from today meeting June 14

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

- Discussion of UEFI HTTPS boot
  - Two stage provisioning could be possible
  - Mutual authentication should be required for multi-tenant environments

- Discussion on distribution of keys for PKI – initial authentication of provisioning server by the DPU
  - The processes for doing this should be an open item for the forum
  - The specifications to not prescribe exactly when/who/how those keys are programmed in the DPU

- Shared document (posted on Github) for SDO and BRSKI discussion

- Discussion – tentative goal proposed to adopt a selection (FDO, sZTP, others) to implement (or select) an open source implementation of a provisioning server

- Discussion on ZTP, used by SONIC
  - Might be a good example of a procedure for private networks.
  - (It was pointed out that it lacks endpoint authentication)
  - Question was asked if there was a way for us to align with the SONIC community on our specs

- Issue of inventory was raised again
  - Provisioned image could be determined by “inventory” or SKU of DPU
  - this could be determined by a device unique id like serial number
  - but it might be better to have a query/response for the inventory for a more robust and flexible solution

- Reiteration of the two markdown pages
  - inventory
  - Provisioning and Discovery

## Lifecycle & Updates
- nothing

## Scheduling
- nothing

## Action items:
- Juniper: present demo for SZTP next week 6/21
- Marvel: update manual provisioning method on MD
