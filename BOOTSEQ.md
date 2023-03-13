# Boot sequencing

## TLDR

This page is high level overview of assumptions, use cases and considerations for coordination between Server (aka Host) and DPU.
The detailed solutinos and options will have their own pages.

For example, the first initial bootup part, is detailed here [COORDINATION.md](./COORDINATION.md) considering in-band via different PCIe mechanisms or OOB via Server BMC assisted mechanisms.

Coordinated shutdowns, reboots, crashes, error handling will be detailed on a separate pages and link will be posted here.

## Terms

| Term | Description
|------|----------------------------------------------------------------------|
| AMC  | Add-in Management Controller on a PCIe card                          |
| BMC  | Baseboard Management Controller                                      |
| CMC  | Chassis Management Controller                                        |
| CRS  | PCI Configuration Request Retry Status                               |
| DPC  | Downstream Port Containment                                          |
| FRU  | Field Replacement Unit                                               |
| IB   | In-Band communications over the PCIe interface                       |
| OOB  | Out-of-Band communication over I2C, NC-SI, UART, etc                 |
| UEFI | Unified Extensible Firmware Interface previously referred to as BIOS |
| xPU  | SmartIO, SmartNIC, DPU, or IPU                                       |


## Entities

Entities optionally participating in the boot/reboot/shutdown processes:

- Platform UEFI/BIOS of the Server
- Platform BMC of the Server
- xPU BMC (also called AMC sometimes)
- xPU ATF/UEFI/BIOS
- additional xPUs that reside within the Server
- Optional Chassis Management Controller (aka CMC)

## Use cases

Following use cases are considered first:

- One or more xPUs in a single server
- A single xPU attached to multiple CPUs within the same server
- The xPU is powered on at the same time as the server BMC using Auxiliary power
- The server BMC controls when the xPU and server CPUs are powered with Primary power
- The xPU boots from the network or on-board non-volatile storage
- xPUs can be with and without a local AMC (aka NIC BMC) on Board

The multi-host case where one or more xPUs are attached to multiple servers within an enclosure is deferred to another exercise.  This requires a Chassis Management Controller (CMC).

## Assumptions

- Resetting the xPU is assumed to cause PCIe errors that will crash the server (at least for some scenarios).
  - This will cause challenges for deploying an OS on the xPU and when running multi-xPU deployments. Perhaps OPI should work with Linux Foundation to encourage host and OS developers to properly manage the response by the host and the xPU(s) from a Reset
  - There are different kinds of resets. Genrally XPU core complex reset and XPU PCIe MAC reset should be separated for minimal disruptive ISSU etc. but a complete SoC reset of the xPU will cause PCIe errors and needs reboot the server.
- There are use cases for network boot only.
- xPU boot can take some time (full linux distro)
- There is communication between some of the entities from above
- While the xPU is powered by 3.3Vaux only
  - OOB communications from the server BMC shall be limited to accessing the FRU (EEPROM) and/or security chip on the xPU if they are available
  - There shall be no IB communication from the server CPU because it is in S5 state (powered off)
- The xPU may ignore transitions in the PERST# signal
- The xPU may require separate mechanisms to reset its PCIe interface and compute complex
- The xPU may start to boot as soon as it detects the stable application of +12V and +3.3V
- The xPU is agnostic as to if +12V and +3V are derived from Primary or Auxiliary power
- The xPU has no dependencies on the status of the server OS
- The server OS may have dependencies on a booted xPU, but may continue to boot when the xPU fails to boot after some appropriate timeout
- Downstream Port Containment (DPC) is not a practical mechanism to hot add PCIe functions during POST
- There are multiple configurations of AMC and BMC
  - BMC and no AMC (will address this use case first)
  - BMC and AMC with BMC as the supervisor
  - BMC and AMC with AMC as the supervisor

## DPU Vendors Survey

- What is the role of a BMC on the xPU regarding startup?
- What casues the ARM cores to start running? Application of 12V? Release of PCIe RESET? Write to a location via NC-SI or I2C? ??
- How does PCIe reset affect the xPU's PCIe interface and the ARM cores?
- What happens to the xPU's PCIe interface when the ARM cores lockup?
- What is the maximum time from release of PCIe RESET until the xPU responds to config cycles?
- How does the xPU determine if it should boot to its OS or some diagnostic FW?
- Is there a difference in xPU PCIe behavior between an OS running or diagnoostic FW?
- What is the method for determining the xPU boot status via PCIe, I2C and NC-SI?

## Bootup

See different approaches for in-band vs oob on bootup ccordination [here](./COORDINATION.md)

- Server is Powered On
- xPU receives power and starts booting
  - These events may not be synchronous. The xPU may be installed in a PCIe slot providing persistent power. xPU power cycle may be a function of the host BMC (and potentially the xPU BMC) rather than the power state of the host
- Should Host OS wait for xPU to finish boot ?
  - For some functions, it must (e.g. NVMe namespace that it needs to boot from)
  - With a pre-provisioned, local host OS present, and for other functions (most notably, networking), the host can continue its boot sequence. Ports should appear down from the host's perspective until the xPU is ready to handle traffic from/to host.
  - Host OS continues to its boot. xPU needs to either respond to PCIe transactions or respond with retry config cycles(CRS) till it is ready. This automatically holds the Host BIOS. Note that Host BIOS may need to be modified to allow longer CRS response time before timing out.
- Should xPU OS wait for host to finish boot ?
  - How tightly coupled should host and xPU be during boot phase?
- who tells Host OS/xPU OS to wait and how ?
  - PCIe has a mechanism for config retry cycles(CRS) that can be used.
- who tells Host OS to to continue / start booting ?
  - These events may not be synchronous. The xPU may be installed in a PICe slot providing persistent power. xPU power cycle may be a function of the host BMC (and potentially the xPU BMC) rather than the power state of the host
- is Host CPU halted ?
- can we wait in UEFI/BIOS in case of network boot via xPU ?
  - Some UEFI PXE/HTTP boots have timeouts, sometimes painfully short ones (as low as 30 seconds)
  - If the XPU provides an UNDI driver for network boot in its OPROM, it can have the Host BIOS wait.
- Should host continue booting even if xPU failed to boot after timeout ?
  - Yes. XPU services/devices will be unavailable
- More...

## xPU and host reboot

- xPU OS/FW reboot assumed to cause Host OS crash
  - Not in all cases. e.g. NVIDIA's design separates the various components, and an ARM complex reboot should have no direct effect on the host OS. Similarly, some vendors implement ISSU processes for certain FW updates where no disruption is observed.
  - This will cause challenges for xPU OS deployment and multi-xPU depolyments. We should encourage graceful reset behaviour from host/OS/xPUs
- How xPU can reboot ?
  - Can notify Host OS that reboot is about to happen?
  - IRQ?
  - OS specific?
  - Through out-of-band link to host BMC?
  - xPU has to reboot without causing complete disrutpion, it can result in packet drops/storage IO timeout etc, but not require host reboot. It can send IRQs which would be vendor specific. Another option is to hotunplug all devices from host, which would be disruptive.
- More...

## Graceful shutdown

- Should Host notify to xPU when reboot requested (linux shutdown/reboot command) and how ?
  - How does xPU reaction change if it is installed in a slot providing persistent power, implying it should be independent of host?
  - Host reboot would cause PCIe PERST# signal to be asserted to xPU. xPU can use that to detect Host reboot.
- Should xPU perfoms any actions on Host reboot request ?
  - Yes. it could be context cleanup etc, the host can be baremetal rented to another tenant after a reboot.
- When/How xPU can notify Host OS that reboot can happen/continue ?
  - Can xPU deny a host reboot/powercycle event?
  - Does persistent power change this behavior?
  - xPU could control its PCIe config/Bar  responses once it sees PERST# signal from host. the Host BIOS will then wait for PCIe config cycles to be succesful during reboot.
- How does host respond to a xPU going off line? (during lifecycle event etc.)
- More...

## Error handling

- what happens during Host OS crash / core dump ?
- this should affect xPU only to the extent that xPU loses connectivity to host. Other services on xPU should continue.
- what happens during xPU OS crash / core dump ?
- There would be permanent disruption of services till xPU is rebooted. this may need host reboot or rescan of PCIe bus.
- does a multi-xPU deployment change anything?
- More...
