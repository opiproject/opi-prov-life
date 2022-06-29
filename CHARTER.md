# Charter

_This is just a bunch of ideas at this moment.
We recognise this is too big to our sub-group to handle.
Once we finish with this list, we shall de-prioritize and/or drop some of those._

## Security

- Both in lifecycle, discovery and installations/provisioning
- Consider roll into one of the below

## Discovery

- Discovery of the DPU and optionally BMC
- DNS / LLDP / SSDP ?
- SDN controller ?
- Discover could be part of the Provisioning bullet
- Discover from
  - Network
  - Host
  - BMC
  - Other…

## Inventory

- Identity - Vendor, SN, P/N, SW/FW versions, …
- Capabilities of the DPU
  - System resources: Core, Memory,  storage,  …
  - Offload / Accelerations: hw vs sw capabilities
    - Compression, Encryption,FPGA, GPU,Raid, …
    - Algorithms, Features, protocols, …
    - Performance capabilities BW, IOPs, Latency…
  - Capabilities are dynamic and can be added/removed (i.e., via FPGA)
  - Consumed capacity/resources vs remaining capacity / resources
- Advertising to ? i.e. SDN controller
- Advertise management end point) for example port) that API will consume
- Persistent High-Available Storage AND State – is it something we need? Or part of the sub-group ?

## Initial provisioning / Initial boot image…

- PXE, local DPU storage boot, boot from the host, boot from BMC
- Zero-trust
- Secure boot and image signing
- Consider making this a priority item of the focus
- UEFI boot, selecting the OS and RootFS is one option

## Boot sequencing & crash recovery

- **IS IT PART OF THE SCOPE? Or just lower priority ?**
- Who boots first : dpu / dpu bmc / host / host bmc ?
- Wait for boot ? control boot order…
- Controlled reset – what is the sequence?
  - Only DPU v only Host vs Both vs all DPUs (multiple DPU in the host)
  - DPU power cycle
- Uncontrolled reset
  - When DPU SW/Kernel crashes
  - When DPU FW/NIC/powerloss crashes or reboots
  - When Host crashes soft reset
  - When Host crashes – power cycle / reset
  - When BMC crashes?
  - PCIe uncorrectable errors
  - **Maybe Uncontrolled reset is lower priority**
- DPU is not booting at all
  - Corrupted flash / corrupted/missing DIMMs
  - Understand reason for not booting – APIs…
  - Ability to debug this state

## Lifecycle & Updates

- DPU fw images
- DPU sw images
- DPU FPGA
- BMC fw
- Host ? – agent/driver on the host talking to DPU
- Non-disruptive upgrades
  - host application is not affected if DPU is updated
  - DPU is not affected if HOST is updated
- Version dependencies
- Return versions as part of inventory/discovery
- New: who initializes the update ?
  - Host / bmc /dpu ?
  - Focus on DPU interface that is exposed/provided for the update
- New: where the update images live?
  - TBD
- Access/user/authentication rights – who is allowed to do what and through which interface

## Deploy/Scheduling services on DPU (as oppose of OS deploy)

- Where services are deployed from? Host & Network
  - According to Security / Access rights
- Ability to Deploy at large scale
- Stacking / Chaining of services
  - More API question and not deploy question?
- Running environment / context
  - Kernel modules
  - Containers
    - Do we care docker or rocket or pods whatever?
    - Though operator vs k8s-like pods / services
  - Apps not in container
  - VMs – out of question
- Lifecycle of services (+upgradability)
  - Containers vs kernel – different lifecycle methods
  - Some services have to be non-disruptive (no down time)
  - Resources, Compatibility, Security
  - BMC involvement in the services lifecycle – yes/no?
  - Ripple / Cascade upgrades with dependencies
- Security & Permissions
  - Signed services (like signed images)
  - Trusted services (covered by signed services) signed authority?
  - Third party vs Official/Published vs Vendor vs Private services
  - Access/Permissions rights to deploy container vs kernel module
  - Selinux / Sandbox / c-groups / network namespace…
- Scheduling
  - Services in random nodes (according to resources and constrains)
  - Services on specific nodes (for example to accommodate specific app on the host)
  - Ripple / Cascade scheduling with dependencies
  - Resources available vs taken of the DPU
    - CPU, memory, network, accelerators, local storage,…
  - Compatibility
    - From DPU Vendor and P/N
    - With other services running already
    - Underlying base distro/linux
- **State** of the services
  - High availability of the state?
  - State across several DPUs
- Migration
  - From Host to DPU and back
  - From DPU to another DPU
  - ...
- Open questions
  - Do we have any services/apps that will require DPU reboot ?
  - ...

## Monitoring & Telemetry

- Centralized and aggregated separated service
  - First inside the DPU locally
  - Across fleet of DPUs (optionally)
- Propagation & subscription model
  - PUB/SUB
  - Pulling
  - Streaming
- Storage to hold all this data
  - Streaming platforms…
  - Permanent storage…
  - External to DPU
  - Internal to DPU
  - If Host server is available host storage option
  - How much data we save back…
- Plugins
  - System syslog plugin (mostly kernel and system)
  - BMC plugin
  - Most of the services are user space and containers
  - Native to other custom services
  - What protocols do we use? gRPC? maybe out of scope ?
- Verbosity
  - Filtering on sender (service) side
  - Filtering on aggregator side
- Classes
  - Logging
    - Is this just type of telemetry or event below ?
    - Telemetry
    - Environmental (temp)
    - State of resources / load (cpu, mem,..)
    - Error counts from offload/fw/hw/…
    - Usage statistics (bw, iops, packets … for billing)
  - Events / Alerts
    - State change
    - Service up/down/new
  - Capabilities
    - See the capabilities bullet…
    - Does this change much ?
- Performance
  - Possible out of scope …
  - Due to limited resources of the DPU
- Telemetry loss
  - Possible out of scope …
  - What happens when DPU loses communication?
  - How to offload / trim / discard the aggregated?
- Implementation questions
  - Can we utilize IPDK framework and start implementing all the things we discuss above?
  - This could be a good vehicle to try those ideas in the real-life application
  - The APIs today in SPDK/IPDK/DPDK are too low level
  - We can start with those frameworks…
  - Action item: bring somebody from IPDK (Dan.D) ?
  - Think about Telegraf , Fluentd and InfluxDB...
