# Inventory

## Local Inventory

- OPI adopted [SMBIOS](https://www.dmtf.org/standards/smbios) for DPUs and IPUs
- SMBIOS is used for local access inside the DPUs and IPUs to get BIOS and System information about DPUs and IPUs
- SMBIOS is a standard way to get similar information from servers, so DPUs and IPUs adoption make sense
- On Linux one can access SMBIOS info locally via `dmidecode` or `/sys/class/dmi/id/`, as an example
- SMBIOS provides those types of information:
  - bios
  - system
  - baseboard
  - chassis
  - processor
  - memory
  - cache
  - connector
  - slot
- OPI will define mandatory tables and fields for OPI compliance.
- OPI will produce smbios-validation-tool like [google](https://github.com/google/smbios-validation-tool)
- Check more info in Provisioning TWG <https://github.com/opiproject/opi-prov-life>

Example from DELL:

```bash
$ dmidecode -t system

Handle 0x0100, DMI type 1, 27 bytes
System Information
        Manufacturer: Dell Inc.
        Product Name: PowerEdge R750
        Version: Not Specified
        Serial Number: 3Z7CMH3
        UUID: 4c4c4544-005a-3710-8043-b3c04f4d4833
        Wake-up Type: Power Switch
        SKU Number: SKU=NotProvided;ModelName=PowerEdge R750
        Family: PowerEdge
```

Example from Nvidia:

```bash
$ dmidecode -t system

Handle 0x0001, DMI type 1, 27 bytes
System Information
        Manufacturer: https://www.mellanox.com
        Product Name: BlueField SoC
        Version: 1.0.0
        Serial Number: Unspecified System Serial Number
        UUID: 2e3bc1d1-e205-4830-a817-968ed1978bac
        Wake-up Type: Power Switch
        SKU Number: Unspecified System SKU
        Family: BlueField
```

Example from Marvell:

```bash
$ dmidecode -t system

Handle 0x0001, DMI type 1, 27 bytes
System Information
        Manufacturer: Marvell
        Product Name: ebb106
        Version: unknown
        Serial Number: unknown
        UUID: ea04f2b5-0b7a-4ba9-9b46-13ea9f5f1b95
        Wake-up Type: Power Switch
        SKU Number: Not specified
        Family: Octeon 10
```

## Remote Inventory

- Redfish
  - what if there is no NIC BMC and no IPU IMC ? Run redfish server on the ARM cores
- gRPC
  - add new gRPC service and RPC calls for inventory query

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
