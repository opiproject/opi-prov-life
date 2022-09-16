# Boot sequencing

## Entities

Entities optionally participating in the boot/reboot/shutdown processes:

- Platform UEFI/BIOS of the Server
- Platform BMC of the Server
- xPU BMC
- xPU ATF/UEFI/BIOS

## Asumptions

- Resetting the xPU is assumed to cause PCIe errors that will crash the server.
- There are use cases for network boot only.
- xPU boot can take some time (full linux distro)
- There is communication between some of the entities from above
- More...

## Bootup

- Server is Powered On
- xPU receives power and starts booting
- Should Host OS wait for xPU to finish boot ?
  - who tells Host OS to wait and how ?
  - who tells Host OS to to continue / start booting ?
  - is Host CPU halted ?
  - can we wait in UEFI/BIOS in case of network boot via xPU ?
- Should host continue booting even if xPU failed to boot after timeout ?
- More...

## DPU reboot

- xPU OS/FW reboot assumed to cause Host OS crash
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
