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
- There are use cases for network boot only.
- xPU boot can take some time (full linux distro)
- There is communication between some of the entities from above
- More...

## Bootup

- Server is Powered On
- xPU receives power and starts booting
- Should Host OS wait for xPU to finish boot ?
  - For some functions, it must (e.g. NVMe namespace that it needs to boot from)
  - With a pre-provisioned, local host OS present, and for other functions (most notably, networking), the host can continue its boot sequence. Ports should appear down from the host's perspective until the xPU is ready to handle traffic from/to host.
- who tells Host OS to wait and how ?
- who tells Host OS to to continue / start booting ?
- is Host CPU halted ?
- can we wait in UEFI/BIOS in case of network boot via xPU ?
  - Some UEFI PXE/HTTP boots have timeouts, sometimes painfully short ones (as low as 30 seconds)
- Should host continue booting even if xPU failed to boot after timeout ?
- More...

## DPU reboot

- xPU OS/FW reboot assumed to cause Host OS crash
  - Not in all cases. e.g. NVIDIA's design separates the various components, and an ARM complex reboot should have no direct effect on the host OS. Similarly, some vendors implement ISSU processes for certain FW updates where no disruption is observed.
- How xPU can reboot ? Can notify Host OS that reboot is about to happen ? IRQ ? OS specific ?
- More...

## Graceful shutdown

- Should Host notify to xPU when reboot requested (linux shutdown/reboot command) and how ?
- Should xPU perfoms any actions on Host reboot request ?
- When/How xPU can notify Host OS that reboot can happen/continue
- More...

## Error handling

- what happens during Host OS crash / core dump ?
- what happens during xPU OS crash / core dump ?
- More...
