# Boot sequencing

## Entities

Entities optionally participating in the boot/reboot/shutdown processes:

- Platform UEFI/BIOS of the host
- Platform BMC of the host
- xPU BMC
- xPU ATF/UEFI/BIOS
- additional xPUs that reside within the host

## Assumptions

- Resetting the xPU is assumed to cause PCIe errors that may impact the operation of the host OS.
  - There are different kinds of resets. Generally XPU core complex reset and XPU PCIe MAC reset should be separated for minimal disruptive ISSU etc. but a complete SoC reset of the xPU will cause PCIe errors and may result in a controlled or uncontrolled host reset.
  - See  _Host Reset or Crash_ and _xPU Reset or Crash_ sections for specifics.

- There are use cases for network boot only where the host is used only to provide power and cooling to one or more xPUs.
- xPU boot can take some time (full linux distro)
- There is communication between some of the entities from above
- More...

## Simultaneous Host and xPU Bootup

1. Host and xPU are simultaneously powered up
2. Host and xPU are simultaneously restarted

How tightly coupled should host and xPU be during boot phase?

- This is likely use-case dependent, so OPI should accomodate all cases from tightly coupled to completely independent operation.

Events:-

- Host is Powered On
- xPU receives power and starts booting
- Should host OS wait for xPU to finish boot ?
  - For some functions, it must (e.g. NVMe namespace that it needs to boot from)
  - With a pre-provisioned, local host OS present, and for other functions (most notably, networking), the host can continue its boot sequence. Ports should appear down from the host's perspective until the xPU is ready to handle traffic from/to host.
  - Host OS continues to its boot. xPU needs to either respond to PCIe transactions or respond with retry config cycles (CRS) till it is ready. This automatically holds the host BIOS.
  - The xPU may hold its PCIe channel down during xPU boot. When the xPU has finished booting, it generates a hot-plug event, and the host will insert the xPU into the PCIe root complex. Is this true for all host OS?

- Who tells host OS/xPU OS to wait and how ?
  - PCIe has a mechanism for config retry cycles (CRS) that can be used
  - Host BIOS may need to be modified to allow longer CRS response time before timing out

- Host BIOS may sequence the start-up either automatically using orchestration built into BIOS or explicitly through commands sent by external orchestrator to BIOS
  - Provide power to xPU without starting host
  - Wait for xPU to complete startup
  - Boot host

- Who tells host OS to continue / start booting ?
  - Default is for host OS continues boot process unless it is paused by CRS signal from xPU ?

- Are there any circumstances under which host CPU is halted ?

- Can we wait in UEFI/BIOS in case of network boot via xPU ?
  - Some UEFI PXE/HTTP boots have timeouts, sometimes painfully short ones (as low as 30 seconds)
  - If the xPU provides an UNDI driver for network boot in its OPROM, it can have the host BIOS wait.

- Should host continue booting even if xPU failed to boot after timeout ?
  - Yes. xPU services/devices will be unavailable
  - Yes. May be needed to overwrite a corrupt image on the xPU

- Should xPU OS wait for host to finish boot ?
  - Only if use case requires the xPU to have dependencies on resources owned by the host ?

- More...

## Host Bootup/Restart

1. xPU already executing in always-on-power slot
2. Host reboot; xPU ignores reboot and continues executing during host reboot

Use cases

1. Host is providing power and cooling to xPU; there is no dependency by the xPU on any host functionality

Events:-

- Host is started/restarted
- All resources provided by the xPU are already online and can be used by host during boot phase
- Host can use the xPU for network boot without timeout issues
- More...

## xPU Bootup/Restart

1. xPU restarting after reset or crash; DPC event did not result in host restart or crash
2. xPU powered on using host BIOS to provide power to PCIe slot without starting host

- xPU starts booting
- xPU completes boot and issues hot-plug event
- Host updates root complex with xPU resources and makes those resources available to services executing on the host
- _Note: This assumes host will properly execute a hot-plug event_
- More...

## Host OS Reset or Crash

- During host OS Reboot or Crashes, the host writes to Chipset Reset Control Register (i.e., IO port 0xCF9) or to the ACPI FADT
- The Platform (CPLD) firmware can monitor for these events and trigger host BMC/BIOS to coordinate with xPU to take appropriate action. Potential actions include:-
  - Ensuring that xPU is gracefully shut down before a full host reboot
  - Force reset xPU after a timeout if xPU has failed to gracefully shutdown
  - Allow xPU to continue executing if there is no host/xPU dependency

## xPU OS Reset or Crash

xPU OS/FW crash and/or independent reboot will result in a Downstream Port Containment (DPC) event on the PCIe bus

- The host (BMC/BIOS) firmware and host OS needs to implement DPC and hot plug requirements defined in PCIe Firmware spec 3.3 for these scenarios to work.

- Coordinated shutdown
  - The xPU can provide in-band or out-of band signalling to host OS that it intends to restart. The host OS can respond by either
    - Ignoring the event if the host OS intends to continue execution after the remove event, containing the subsequent surprise removal
    - Host OS reset (see above), to coordinate reboot of other xPUs on the host bus

- xPU OS crash or reset.
  - xPU will trigger DPC/CI surprise remove event.
  - Host OS and BMC/BIOS will need to handle these events gracefully so that host OS remains active for long enough for the xPU OS crash dump to complete.
  - If host OS is tightly coupled to xPU function, it can elect to gracefully shutdown/reboot when the xPU crashdump has completed.
  - Alternatively, if host OS is losely coupled or decoupled from xPU (if xPU is acting as a bump in the wire service for example), the host can
    - Elect to contain the event, removing the xPU from the bus table. When the xPU restarts, it will generate a hot-plug event on the PCIe bus, allowing the host BIOS/OS to reinsert the xPU.
    - Note: This behavior is currently a largely untested feature in Linux host OS, and is not available on Windows

  - _Further discussion required_ NVIDIA's design separates the various components, and an ARM complex reboot should have no direct effect on the host OS. Similarly, some vendors implement ISSU processes for certain FW updates where no disruption is observed.
  - This will cause challenges for xPU OS deployment and multi-xPU depolyments. We should encourage graceful reset behaviour from host/OS/xPUs

- How does xPU notify host OS that reboot is about to happen?
  - IRQ?
  - OS specific?
  - Through out-of-band link to host BMC?
  - xPU has to reboot without causing complete disrutpion, it can result in packet drops/storage IO timeout etc, but not require host reboot. It can send IRQs which would be vendor specific. Another option is to hotunplug all devices from host, which would be disruptive.

- More...

## Graceful shutdown

- Should host notify to xPU when reboot requested (linux shutdown/reboot command) and how ?
  - How does xPU reaction change if it is installed in a slot providing persistent power, implying it should be independent of host?
  - Host reboot would cause PCIe PERST# signal to be asserted to xPU. xPU can use that to detect host reboot.

- Should xPU perfoms any actions on host reboot request ?
  - Yes. it could be context cleanup etc, the host can be baremetal rented to another tenant after a reboot.

- When/How xPU can notify host OS that reboot can happen/continue ?
  - Can xPU deny a host reboot/powercycle event?
  - Does persistent power change this behavior?
  - xPU could control its PCIe config/Bar  responses once it sees PERST# signal from host. the host BIOS will then wait for PCIe config cycles to be succesful during reboot.

- How does host respond to a xPU going off line? (during lifecycle event etc.)

- More...

## xPU OS Installation mode

- An xPU may go through multiple reboots during xPU OS installation. Consequently, the host OS may choose to bring the xPU and host PCIe link down prior to xPU OS install, and to hold the link down until the install process has completed
  - The xPU can use the NC-SI connection to signal to the host that xPU OS installation has completed, allowing the host to enable the host/xPU PCIe link
  - Alternatively, if the host/xPU has no NC-SI connection, then the host can wait for a pre-defined timeout period before enabling the PCIe host/xPU link
  - The installer executing on the host, may be able to take down and restore the host/xPU link without explicit support for DPC in BMC/BIOS

## Multi-xPU install

- does a multi-xPU deployment change anything?
- More...
