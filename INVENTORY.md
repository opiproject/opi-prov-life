# Inventory

## Local Inventory

Moved to [here](./inventory/LOCAL.md)

## Remote Network Inventory

- Redfish
  - what if there is no NIC BMC and no IPU IMC ? Run redfish server on the ARM cores
- gRPC
  - add new gRPC service and RPC calls for inventory query
  - see proto definition [here](https://github.com/opiproject/opi-api/blob/main/common/v1/inventory.proto)
  - see example implementation [here](https://github.com/opiproject/opi-smbios-bridge)

## Remote Host/BMC Inventory

Host host/platform BMC access inventory information from the DPU/IPU ?

- NC-SI
- PLDM
- VPD
- other...

VPD Example from Nvidia:

```bash
$ lspci -s ca:00.0 -vvv | grep -A 11 "Vital Product Data"
        Capabilities: [48] Vital Product Data
                Product Name: BlueField-2 E-Series SmartNIC 100GbE/EDR VPI Dual-Port QSFP56, PCIe Gen4 x16, Crypto Enabled, 16GB on-board DDR, FHHL
                Read-only fields:
                        [PN] Part number: MBF2M516A-EEEOT
                        [EC] Engineering changes: A1
                        [V2] Vendor specific: MBF2M516A-EEEOT
                        [SN] Serial number: MT2022X19080
                        [V3] Vendor specific: 805647c144a9ea1180000c42a198b662
                        [VA] Vendor specific: MLX:MN=MLNX:CSKU=V2:UUID=V3:PCI=V0:MODL=BF2M526A
                        [V0] Vendor specific: PCIeGen4 x16
                        [RV] Reserved: checksum good, 1 byte(s) reserved
                End
```

## Leftovers

this page talks about common inventory information information provided by DPU vendors in a same format

- Questions:
  - Where does it run?
    - service/container running on ARM cores?
    - DPU firmware?
    - DPU Bootloaders?
    - DPU BMC?
    - External, as a control plane component?
  - Is it always available/running?
    - if this is e.g. a UEFI application then it would only be available pre-boot.
  - What information does it provide ?
    - Vendor, SN/PN, ...
    - Credentials
    - Virtual location in the host (mother board slot or PCIe BDF?)
    - Capabilities (cou, mem, offloads,...)
    - What protocols and/or privisoning methods are supported
    - What DPU belongs to what Host ?
  - Does it need external sources of information to perform its function?
    - mapping to a host / host PCI topology is probably not possible from the perspective of a PCI device.
  - What protocol and what format is used?
    - GraphQL? Rest ? gRPC?
    - JSON/XML/Protobuf?
    - in FW we can't use TCP probably, so look into ICMP
  - Can we simplify the approach?
    - If we break this into a data collector application and a data store/query service as two separate entities:
      - Their actual run times can be non contemporaneous (e.g. first one runs periodically, or when things change, or remotely, ...).
      - Also a benefit exists here where the latter can be a generic OPI software, while the former is vendor-specific by nature.

this is not the end... just pausing here the doc...
