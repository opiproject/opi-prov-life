# Remote Network Inventory

This is designed for external control plane that wishes to access inventory, like serial number, remotely.
Also, potentially, also match what xPU resides in what Server.

see [Terms](../boot/README.md#terms)

## Redfish

Advantage: new standard for servers today, has big ecosystem around it

- what subset of redfish to implement on xPU BMC (aka AMC) ?
- what we re-use and/or contribute back to [OpenBMC](https://github.com/openbmc)?
- what if there is no xPU BMC (aka AMC) ? Run redfish server on the ARM cores?
- How [Redfish SmartNIC White Paper](https://www.dmtf.org/sites/default/files/standards/documents/DSP2063_1.0.0.pdf) is related ?
- Can we create [Mockups](https://github.com/DMTF/Redfish-Mockup-Creator) for DPUs?
- See DMTF ecosystem around resfish [here](https://github.com/search?q=org:DMTF+redfish&type=repositories)

### Redfish Vendor examples

Example from DELL:

```bash
$ curl -qkL -u root:password https://10.240.76.127/redfish/v1/UpdateService/FirmwareInventory/Installed-0-24.35.20.00
{
  "@odata.context": "/redfish/v1/$metadata#SoftwareInventory.SoftwareInventory",
  "@odata.id": "/redfish/v1/UpdateService/FirmwareInventory/Installed-0-24.35.20.00",
  "@odata.type": "#SoftwareInventory.v1_2_1.SoftwareInventory",
  "Description": "Represents Firmware Inventory",
  "Id": "Installed-0-24.35.20.00",
  "Name": "Mellanox Network Adapter - B8:CE:F6:CC:6A:16",
  "ReleaseDate": "00:00:00Z",
  "SoftwareId": "0",
  "Status": {
    "Health": "OK",
    "State": "Enabled"
  },
  "Updateable": true,
  "Version": "24.35.20.00"
}
```

Example from Nvidia:

```bash
$ curl -qkL -u root:password https://bf2-bmc-ip/redfish/v1/UpdateService/FirmwareInventory/e4013534_running
{
    "@odata.id": "/redfish/v1/UpdateService/FirmwareInventory/e4013534_running",
    "@odata.type": "#SoftwareInventory.v1_1_0.SoftwareInventory",
    "Description": "BMC image",
    "Id": "e4013534_running",
    "Members@odata.count": 1,
    "Name": "Software Inventory",
    "RelatedItem": [
        {
            "@odata.id": "/redfish/v1/Managers/bmc"
        }
    ],
    "Status": {
        "Health": "OK",
        "HealthRollup": "OK",
        "State": "Enabled"
    },
    "Updateable": true,
    "Version": "2.8.2-20-gfc1389898"
}
```

Example from Marvell:
```bash
tbd
```

Example from Intel:
```bash
tbd
```

Example from AMD:
```bash
tbd
```

## gRPC

Advantage: all modern control planes today use gRPC and not REST

- add new gRPC service and RPC calls for inventory query
- see proto definition [here](https://github.com/opiproject/opi-api/blob/main/common/v1/inventory.proto)
- see example implementation [here](https://github.com/opiproject/opi-smbios-bridge)

### gRPC Vendor examples

- [Nvidia](https://github.com/opiproject/opi-smbios-bridge#nvidia-example)
- [Marvell](https://github.com/opiproject/opi-smbios-bridge#marvell-example)
- [Intel](https://github.com/opiproject/opi-smbios-bridge#intel-example)
- [AMD](https://github.com/opiproject/opi-smbios-bridge#amd-example)
