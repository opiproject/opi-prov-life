# Local Inventory

This is designed for applications running on the xPU to access inventory locally, like serial number.

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

## Validation

In order to test compliance and help DPU vendors to fix errors and missing tables and fields, we adopting validation tools:
- See <https://github.com/opiproject/smbios-validation-tool>
- See <https://wiki.ubuntu.com/FirmwareTestSuite>

## Examples

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

Example from Intel:

```bash
$ dmidecode -t system

Handle 0x0001, DMI type 1, 27 bytes
System Information
        Manufacturer: Intel
        Product Name: Intel(R) Infrastructure Processing Unit (Intel(R) IPU)
        Version: A0
        Serial Number: -------------
        UUID: 8a95d198-7f46-11e5-bf8b-08002704d48e
        Wake-up Type: Power Switch
        SKU Number: -------------
        Family: -------------
```
