# Boot sequencing

## Entities

Entities optionally participating in the boot/reboot/shutdown processes:

- Platform UEFI/BIOS of the Server
- Platform BMC of the Server
- xPU BMC
- xPU ATF/UEFI/BIOS
- additional xPUs that reside within the Server

## Assumptions

- Resetting the xPU is assumed to cause PCIe errors that will crash the server (at least for some scenarios).
  - This will cause challenges for deploying an OS on the xPU and when running multi-xPU deployments. Perhaps OPI should work with Linux Foundation to encourage host and OS developers to properly manage the response by the host and the xPU(s) from a Reset
  - There are different kinds of resets. Genrally XPU core complex reset and XPU PCIe MAC reset should be separated for minimal disruptive ISSU etc. but a complete SoC reset of the xPU will cause PCIe errors and needs reboot the server.
- There are use cases for network boot only.
- xPU boot can take some time (full linux distro)
- There is communication between some of the entities from above
- More...

## Bootup

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
