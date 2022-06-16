# Inventory

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
