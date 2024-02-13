# OPI Device Provisioning & Lifecycle

For the purposes of this specification a 'device' is a PCI device that intends to be compatible with this specification (e.g., a DPU or an IPU).

## PCI Baseline

- **PCIe Endpoint**: Device shall have one or more PCIe interfaces capable of connecting to a PCI root port (RP) as a PCI endpoint (EP).
- **PCIe Compliant**: Device shall be compatible to the PCIe specifications for endpoint (EP) operation.
- **PCIe Enumeration**: Device shall present a fixed set of PCIe devices on each PCI bus during enumeration, and only modify the set of PCIe devices using mechanisms defined within the PCIe specification.
- **Host Boot Option**: Device should offer bootable PCIe devices (e.g., network interface or a bootable block device) on the PCIe bus for use by the attached Host for booting purposes. These PCIe devices shall be compatible with UEFI running on the Host in all boot ordering scenarios (Host boots first, Device boots first, Host and Device boot at the same time).

## ARM Firmware & Security Baseline

- DPU/IPUs with Arm-based processors are required to be compliant with the following specifications. Requirements are to be verified using SystemReady ES ACS v1.2.0 (or newer).
  - BSA ver 1.0c (or newer).
  - SBBR recipe in BBR v1.0 (or newer), which includes implementing UEFI, ACPI, and SMBIOS.
  - BBSR v1.2 (or newer).
- DPU/IPUs with Arm-based processors are recommended to achieve Arm SystemReady ES certification with the additional Security Interface Extension (SIE) certification.

## BMC Connectivity Baseline

1. **PLDM**: Device shall support DMTF Platform Level Data Model (PLDM) type 2 (Monitoring & Control) using a standard PCIe connector physical medium MCTP transport (e.g. SMBus, PCIe VDM, USB). This includes the ability to read and monitor in addition to being able to set thresholds. PLDM is used to support get the following information from the device:
    1. **Temperature**: Temperature sensors shall be supported
    2. **Power**: Voltage sensors shall be supported
    3. **Identification (via FRU)**: The Following fields are defined by [DSP0257](https://www.dmtf.org/sites/default/files/standards/documents/DSP0257_1.0.0.pdf) - 212 (General FRU Record):
        1. **Type, Model, Part Number**: Used to identify the type, model, and part number for this device
        2. **Serial Number**: Used to identify this particular device
        3. **Manufacturer, Manufacturer Date, Vendor, Name, SKU**: Used to identfy the manufacturer including the SKU
        4. **Version**: The hardware version of this device
        5. **Asset Tag**: The same information provided by a QR code on the physical device.

## Life Cycle Management

- **Provisioning**: Device shall support initial provisioning and re-provisioning using sZTP as defined by OPI here.
- **Telemetry**: Devices shall support runtime telemetry using OTEL as defined by OPI here.
- **System Information**: Devices shall support SMBIOS for system information from the device's operating system, as defined by OPI here.
- **Security Keys**: SSH and TLS keys can be loaded onto the device using standard methods used by Linux, or pre-loaded during provisioning.
- **Runtime Configuration**: Devices shall support the capability of reading and writing runtime configuration over the network using an RPC (e.g., Redfish or OpenConfig) secured using TLS. The configuration space is TBD.
- **Software Update**: Devices shall support the capability of updating the software on the device over the network using an RPC (e.g., Redfish or OpenConfig) secured using TLS.
