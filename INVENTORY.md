# OPI DPU/IPU Inventory

This document defines the OPI inventory data model for DPU/IPU devices,
inspired by [OpenStack Ironic](https://docs.openstack.org/ironic/latest/)
inventory format and adapted for DPU/IPU specific needs.

## Overview

OPI supports three ways to access inventory information from a DPU/IPU:
- [Local Inventory](inventory/LOCAL.md) - using SMBIOS
- [Remote Network Inventory](inventory/REMOTE.md) - using Redfish or gRPC
- [Remote Host/BMC Inventory](inventory/HOST.md) - using NC-SI, PLDM, VPD

## OPI Inventory Data Model

Below is an example OPI inventory JSON for a DPU/IPU device:

```json
{
   "cpu_arch": "aarch64",
   "macs": [
      "00:1a:2b:3c:4d:5e"
   ],
   "interfaces": {
      "eth0": {
         "ip": "192.168.1.100",
         "mac": "00:1a:2b:3c:4d:5e",
         "pxe": true,
         "client_id": null
      }
   },
   "cpus": 16,
   "memory_mb": 16384,
   "inventory": {
      "bmc_address": "192.168.2.100",
      "interfaces": [
         {
            "name": "eth0",
            "mac_address": "00:1a:2b:3c:4d:5e",
            "vendor": "0x15b3",
            "product": "0xa2d6",
            "has_carrier": true,
            "ipv4_address": "192.168.1.100",
            "lldp": null,
            "client_id": null
         },
         {
            "name": "eth1",
            "mac_address": "00:1a:2b:3c:4d:5f",
            "vendor": "0x15b3",
            "product": "0xa2d6",
            "has_carrier": true,
            "ipv4_address": "192.168.1.101",
            "lldp": null,
            "client_id": null
         }
      ],
      "disks": [
         {
            "name": "/dev/mmcblk0",
            "vendor": "Samsung",
            "model": "eMMC",
            "serial": "0x12345678",
            "size": 32013852672,
            "rotational": false,
            "wwn": null
         }
      ],
      "boot": {
         "current_boot_mode": "uefi",
         "pxe_interface": "00:1a:2b:3c:4d:5e"
      },
      "system_vendor": {
         "serial_number": "MT2022X19080",
         "product_name": "BlueField-2 DPU",
         "manufacturer": "NVIDIA"
      },
      "memory": {
         "physical_mb": 16384,
         "total": 17179869184
      },
      "cpu": {
         "count": 16,
         "frequency": "2500.000",
         "flags": [
            "fp",
            "asimd",
            "aes",
            "pmull",
            "sha1",
            "sha2",
            "crc32"
         ],
         "architecture": "aarch64"
      },
      "xpu": {
         "vendor_id": "0x15b3",
         "device_id": "0xa2d6",
         "part_number": "MBF2M516A-EEEOT",
         "serial_number": "MT2022X19080",
         "firmware_version": "24.35.20.00",
         "pcie_bdf": "ca:00.0"
      }
   },
   "error": null
}
```

## Fields Description

### Top Level Fields
- **cpu_arch**: CPU architecture of the DPU/IPU (e.g. aarch64, x86_64)
- **macs**: List of MAC addresses of the DPU/IPU network interfaces
- **interfaces**: Network interfaces available for provisioning
- **cpus**: Number of CPU cores on the DPU/IPU
- **memory_mb**: Total memory in MB on the DPU/IPU

### Inventory Fields
- **bmc_address**: IP address of the BMC (aka AMC) on the DPU/IPU
- **interfaces**: Detailed network interface information
- **disks**: Storage devices available on the DPU/IPU
- **boot**: Current boot mode and PXE interface
- **system_vendor**: Vendor information from SMBIOS
- **memory**: Memory information
- **cpu**: CPU information including architecture and flags
- **xpu**: DPU/IPU specific fields including vendor ID, part number,
  serial number, firmware version and PCIe BDF

## References
- [Local Inventory](inventory/LOCAL.md)
- [Remote Network Inventory](inventory/REMOTE.md)
- [Remote Host/BMC Inventory](inventory/HOST.md)
- [OpenStack Ironic Inventory](https://docs.openstack.org/ironic/latest/)