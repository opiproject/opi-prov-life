## Remote Network Inventory

- Redfish
  - what if there is no NIC BMC and no IPU IMC ? Run redfish server on the ARM cores
- gRPC
  - add new gRPC service and RPC calls for inventory query
  - see proto definition [here](https://github.com/opiproject/opi-api/blob/main/common/v1/inventory.proto)
  - see example implementation [here](https://github.com/opiproject/opi-smbios-bridge)
