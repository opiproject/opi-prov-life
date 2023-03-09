# Bootup coordination approaches

Reference [BOOTSEQ.md](./BOOTSEQ.md)

## TLDR

This page is taking assumptions and use cases from [BOOTSEQ.md](./BOOTSEQ.md) page and tries to deep dive on 3 possible solutions of the only initial bootup part.
Two possible solutions are either in-band via different PCIe mechanisms or OOB via BMC assisted mechanisms.
Initial bootup part covers: Server is Powered On, DPU receives power and starts booting, Host OS should wait for DPU to finish booting.

Coordinated shutdowns, reboots, crashes, error handling will be details on a separate page.

1. In-band PCIe

- using CRS approach
- what problems we have today with PCIe, timeouts, errors, retries,...

2. Driver Ready Check

The race condition to consider is when the Host is first to boot before the infrastrastructure is ready to serve to the Host its boot image (from a port or a disk hanging off the infrastructure device (DPU/IPU)).  To ensure that the port / disk is ready to be used by the Host, the driver running in the UEFI / BIOS should check that the infrastructure is ready before trying to PXE boot or to read the boot disk.
Infrastructure devices acting as either a Host peripheral or as an independent entity will benefit from this 'ready' check.

### Virtio-net

The virtio-net device presents its driver in an option ROM (OROM) for UEFI / BIOS.  This driver will stall the PXE boot process until the infrastructure backend is ready.

### Virtio-blk

The virtio-blk device presents its driver in an option ROM (OROM) for UEFI / BIOS.  This driver will stall the requests on the disk until the infrastructure backend is ready.

### NVMe

The NVMe device driver will poll the CSTS.rdy bit to ensure that infrastructure backend is ready before reading or writing.
 
3. Out-band via platform BMC

## Diagram

![OOB Plat BMC Boot coordination power on seq](architecture/OOB-Plat-BMC-Boot-coordination-power-on-seq.png)

### PLDM State sensors

Useful State definitions

| Set ID 129 Software Termination Status        | Status related to firmware of the operating system.                   | Notes/Usage |
| :-----                                        | :-----                                                                | :-----      |
| 1 – Normal                                    | Software termination is not detected.                                 |             |
| 2 – Software Termination Detected             | Software termination is detected.                                     |             |
| 3 – Critical Stop during  Load/Initialization | The software entity failed during loading or initialization.          |             |
| 4 – Run-time Critical Stop                    | The software entity incurred a run-time failure.                      |             |
| 5 – Graceful Shutdown Requested               | The software entity has been requested to shut down gracefully.       |             |
| 6 – Graceful Restart Requested                | The software entity has been requested to restart gracefully.         |             |
| 7 – Graceful Shutdown                         | The software entity has been shut down gracefully.                    |  Delayed power off.  Wait for xPU to halt before removing power           |
| 8 – Termination Request Failed                | The request to terminate the execution of the software entity failed. |             |

| Set ID 192 Boot/Restart Cause | Represents the stimulus that booted the entity.                                                                                                                                                | Notes/Usage |
| :-----                        | :-----                                                                                                                                                                                         | :-----      |
| 1 – Powered Up                | A start of the system is initiated by changing the entity’s state from  powered off to powered on.                                                                                             |             |
| 2 – Hard Reset                | A restart of the system is accomplished by activating the entity’s reset circuitry.                                                                                                            |             |
| 3 – Warm Reset                | A restart of the system is performed by software that does not involve powering the system off or activating the entity’s reset circuitry.                                                     |             |
| 4 – Manual Hard Reset         | A restart is initiated by the user activation of a mechanical device (for example, pressing a button) and bypasses runtime software.                                                           |             |
| 5 – Manual Warm Reset         | A restart is initiated by the user activation of a mechanical device (for example, pressing a button) and does not involve powering the entity off or activating the system’s reset circuitry. |             |
| 6 – System Restart            | A restart of the entity is initiated by entity hardware components and accomplished by activating the system’s reset circuitry.                                                                |             |
| 7 – Watchdog Timeout          | A restart of the entity is initiated in response to a detected system hang condition.                                                                                                          |             |

| Set ID 196 Boot Progress                                                   | System firmware or software booting status.                                               | Notes/Usage   |
| :------------------------                                                   | :-------------                                                                             | :------- |
| 1 – Boot Not Active                                                        | Boot-up of the firmware or software is not active. It may be already functional.          |         |
| 2 – Boot Completed                                                         | The boot process of the firmware or software has completed.                               |         |
| 3 – Memory Initialization                                                  | The boot process is currently initializing the memory.                                    |         |
| 4 – Hard-Disk Initialization                                               | The boot process is currently initializing the hard disk.                                 |         |
| 5 – Secondary Processor(s)Initialization                                   | The boot process is currently initializing the secondary processors.                      |         |
| 6 – User Authentication                                                    | The boot process is processing the user authentication.                                   |         |
| 7 – User-Initiated System Setup                                            | System firmware or BIOS has entered the user system firmware or BIOS configuration setup. |         |
| 8 – USB Resource Configuration                                             | System firmware or BIOS is currently configuring the USB resource.                        |         |
| 9 – PCI Resource Configuration                                             | System firmware or BIOS is configuring the PCI resources.                                 |         |
| 10 – Option ROM Initialization                                             | The option ROM is being initialized.                                                      |         |
| 11 – Video Initialization                                                  | The video controller is being initialized.                                                |         |
| 12 – Cache Initialization                                                  | The cache memory is being initialized.                                                    |         |
| 13 – SM Bus Initialization                                                 | The system firmware or BIOS is initializing the SM Bus.                                   |         |
| 14 – Keyboard Controller Initialization                                    | The system firmware or BIOS is initializing the keyboard controller.                      |         |
| 15 – Embedded Controller/Management Controller Initialization              | The system firmware or BIOS is initializing the embedded management controller.           |         |
| 16 – Docking Station Attachment                                            | The main system unit is attaching to the docking station.                                 |         |
| 17 – Enabling Docking Station                                              | The system firmware or BIOS is enabling the docking station.                              |         |
| 18 – Docking Station Ejection                                              | The main system unit is ejected from the docking station.                                 |         |
| 19 – Disabling Docking Station                                             | The system firmware or BIOS is disabling the docking station.                             |         |
| 20 – Calling Operating System Wake-Up Vector                               | The system firmware or BIOS is starting the operating system.                             |         |
| 21 – Starting Operating System Boot Process (for example, calling INT 19h) | The system firmware or BIOS is booting the operating system.                              |         |
| 22 – Baseboard or Motherboard Initialization                               | The BIOS is initializing the motherboard.                                                 |         |
| 23 – Floppy Initialization                                                 | The BIOS is initializing the floppy drive.                                                |         |
| 24 – Keyboard Test                                                         | The BIOS is testing the keyboard.                                                         |         |
| 25 – Pointing Device Test                                                  | The BIOS is testing the pointing device.                                                  |         |
| 26 – Primary Processor Initialization                                      | The BIOS is initializing the primary processor.                                           |         |

Reference:

- <https://www.dmtf.org/dsp/DSP0249> DMTF DSP0249 State Set Specification
- <https://www.dmtf.org/dsp/DSP0248> DMTF DSP0248 PLDM for Platform Monitoring and Control Specification

### PLDM RDE

Reference:

- <https://www.dmtf.org/dsp/DSP0218> DMTF DSP0218 PLDM for Redfish Device Enablement

### i2c

I2C on it's own is not useful unless there is a protocol defined.  

### NC-SI OEM

### usb

Could use PLDM State Sensor over PLDM over MTCP over USB
Reference:

- <https://www.dmtf.org/sites/default/files/standards/documents/DSP0283_0.1.5WIP10.pdf> (WIP) DMTF MCTP over USB Binding Specification

### others

## Out-of-band xPU BMC
