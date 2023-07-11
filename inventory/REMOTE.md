# Remote Network Inventory

This is designed for external control plane that wishes to access inventory, like serial number, remotely.
Also, potentially, also match what xPU resides in what Server.

see [Terms](../boot/README.md#terms)

## Redfish

- what subset of redfish to implement on xPU BMC (aka AMC) ?
- what we re-use and/or contribute back to [OpenBMC](https://github.com/openbmc)?
- what if there is no xPU BMC (aka AMC) ? Run redfish server on the ARM cores?
- How [Redfish SmartNIC White Paper](https://www.dmtf.org/sites/default/files/standards/documents/DSP2063_1.0.0.pdf) is related ?

## gRPC

- add new gRPC service and RPC calls for inventory query
- see proto definition [here](https://github.com/opiproject/opi-api/blob/main/common/v1/inventory.proto)
- see example implementation [here](https://github.com/opiproject/opi-smbios-bridge)
