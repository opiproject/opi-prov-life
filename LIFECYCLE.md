# Lifecycle Management

## Device Lifecycle scope

High level requirements on a DPU/IPU or subsequently a "device"

* a way to reboot a device
  * see [Boot Sequence](BOOTSEQ.md)
* a way to update a device
  * firmware update (Interfaces to DPU {gRPC, Redfish to AMC, boot to Net})
    * single or multiple FW? (AMC FW seperate)
    * Inband through host driver or OOB through xMC?
    * Inband through DPU core driver?
    * Update implies inventory
    * Redfish prefered by baremetal providers?
    * Validated FW versions?
    * SW Solution provider vs. Bare Metal provider vs. Service Provider (Hyprescallers are Out of Scope)
  * OS update
    *  different than provisioning
    *  ansible (apt, yum) - package updates?
    *  SONiC like monolithic?
  * software/application update
* a way to recover a device
  * reset to a known good state, e.g., factory reset
  * via network - required
  * via host - optional
* a way to debug a device
  * see [Monitoring](MONITORING.md)
  * uniform method of failure data collection and recovery
