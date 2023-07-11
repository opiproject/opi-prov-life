# Inventory

## Local Inventory

This is designed for applications running on the xPU to access inventory locally, like serial number.
Moved to [here](./inventory/LOCAL.md)

## Remote Network Inventory

This is designed for external control plane that wishes to access inventory, like serial number, remotely.
Also, potentially, also match what xPU resides in what Server.
Moved to [here](./inventory/REMOTE.md)

## Remote Host/BMC Inventory

This is designed for HOST, ether CPU or platform BMC to access inventory, like serial number.
As an example, today some information are via FRU for regular PCIe NICs.
Moved to [here](./inventory/HOST.md)

## Open Questions

- What inventory information we mandate in OPI ?
  - Vendor ID, Serial Number, Part Number, ...
  - Credentials to log into the device
  - Virtual location in the host (mother board slot or PCIe BDF?)
  - Capabilities (CPU, Memory, Networking, Compression/Encryption, Other Offloads,...)
  - What protocols and/or privisoning methods are supported
  - What DPU belongs to what Host
- Where does this inventory service runs ?
  - service/container running on main ARM cores?
  - DPU/NIC firmware?
  - DPU Bootloaders?
  - DPU BMC (aka AMC)?
  - External, as a control plane component?
- Is inventory service always available/running?
  - if this is e.g. a UEFI application then it would only be available pre-boot.
- Does inventory service need external sources of information to perform its function?
  - mapping to a host / host PCI topology is probably not possible from the perspective of a PCI device.
- What protocol and what data format is used to fetch inventory from the device?
  - GraphQL? Rest ? gRPC?
  - JSON/XML/Protobuf?
  - in FW we can't use TCP probably, so look into ICMP
- Can we simplify the approach?
  - If we break this into a data collector application and a data store/query service as two separate entities:
    - Their actual run times can be non contemporaneous (e.g. first one runs periodically, or when things change, or remotely, ...).
    - Also a benefit exists here where the latter can be a generic OPI software, while the former is vendor-specific by nature.
