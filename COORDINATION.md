# Bootup coordination approaches

Reference [BOOTSEQ.md](./BOOTSEQ.md)

## TLDR

This page is taking assumptions and use cases from [BOOTSEQ.md](./BOOTSEQ.md) page and tries to deep dive on 3 possible solutions of the only initial bootup part.
Three possible solutions are either in-band via different PCIe mechanisms or option ROM (OROM) for UEFI / BIOS or OOB via BMC assisted mechanisms.
Initial bootup part covers: Server is Powered On, DPU receives power and starts booting, Host OS should wait for DPU to finish booting.

Coordinated shutdowns, reboots, crashes, error handling will be details on a separate page.

## Terms

| Term                   | Description                                                          |
|------------------------|----------------------------------------------------------------------|
| Platform BMC           | Baseboard Management Controller                                      |
| CMC                    | Chassis Management Controller                                        |
| CRS                    | PCI Configuration Request Retry Status                               |
| DPC                    | Downstream Port Containment                                          |
| FRU                    | Field Replacement Unit                                               |
| OOB                    | Out-of-Band communication over I2C, NC-SI, UART, etc                 |
| UEFI                   | Unified Extensible Firmware Interface previously referred to as BIOS |
| xPU                    | SmartIO, SmartNIC, DPU, or IPU                                       |
| xPU BMC, aBMC, AMC     | Add-in Management Controller on a PCIe card                          |
| xPU Mgmt               | Terminates OOB Mgmt from Platform BMC, maybe xPU BMC or other MCU    |

## 1: In-band PCIe

In-Band refers to PCIe config access to the xPU from UEFI running on the server cores.

1. All xPU devices shall have a standard PCI PF0 interface (PCI Physical Function 0) that is fully compliant to the PCIe specification
   - The PF0 interface shall be functional 3 seconds after PERST# is de-asserted
   - The PF0 interface shall be functional when the compute complex on the xPU is running, locked up or halted
2. Define a new PCI Class Code and corresponding PCI architected Extended Capability Structure for xPUs
   - Define PCI sub-classes for xPUs with and without an AMC
   - The PCI architected Extended Capability Structure shall include items 3-7 below
3. Define a read/write PF0 OS_BOOT_SELECT register with the following bits defined
   - Network OS boot (bit only - does not include the actual path and credentials)
   - OS on xPU’s non-volatile storage
   - UEFI on xPU’s non-volatile storage
   - Maintenance OS/FW on xPU’s non-volatile storage
   - Other
4. Define a read only PF0 OS_STATUS register with the following bits defined
   - Not started
   - Booting
   - Booted
   - Stalled or locked-up
   - Halted
5. Define a read only PF0 CRASHDUMP_STATUS register with the following bits defined
   - Not started
   - In progress
   - Complete
6. Define a read only PF0 MAX_BOOT_TIME register with a range from 0 to 1600 seconds
7. Define a read/write PF0 RESET register with individual bits to resets specific segments of the xPU
   - PCI interface
   - CPU complex
   - AMC
   - Accelerator1
   - Accelerator2
   - other

## 2: Driver Ready Check

The race condition to consider is when the Host is first to boot before the infrastrastructure is ready to serve to the Host its boot image (from a port or a disk hanging off the infrastructure device (DPU/IPU)).  To ensure that the port / disk is ready to be used by the Host, the driver running in the UEFI / BIOS should check that the infrastructure is ready before trying to PXE boot or to read the boot disk.
Infrastructure devices acting as either a Host peripheral or as an independent entity will benefit from this 'ready' check. This ready check enables the common case of shared power across DPU/IPU and server(s), and enables parallel startup of the DPU/IPU and attached host(s) once shared power is applied.

### Assumptions

The infrastructure devices must be ready on the PCIe bus(es) at enumeration time which allows the driver to address the race condition.

### Virtio-net

The virtio-net device presents its driver in an option ROM (OROM) for UEFI / BIOS.  This driver will stall the PXE boot process until the infrastructure backend is ready (via a driver specific signaling).

### Virtio-blk

The virtio-blk device presents its driver in an option ROM (OROM) for UEFI / BIOS.  This driver will stall the requests on the disk until the infrastructure backend is ready (via a driver specific signaling).

### NVMe

The NVMe device driver will poll the CSTS.rdy bit to ensure that infrastructure backend is ready before reading or writing.

## 3: Out-band via platform BMC

### Diagram

![OOB Plat BMC Boot coordination power on seq](architecture/OOB-Plat-BMC-Boot-coordination-power-on-seq.png)

### PLDM State sensors - PLDM

Note: PLDM is over MTCP which can be caried over i2C, serial, USB and other physical connections

Useful State definitions

| Set ID 129 Software Termination Status        | Status related to firmware of the operating system.                   | Notes/Usage |
| :-----                                        | :-----                                                                | :-----      |
| 0 – Unknown                                   | Unknown                                                               |             |
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
| 0 – Unknown                   | Unknown                                                                                                                                                                                        |             |
| 1 – Powered Up                | A start of the system is initiated by changing the entity’s state from  powered off to powered on.                                                                                             |             |
| 2 – Hard Reset                | A restart of the system is accomplished by activating the entity’s reset circuitry.                                                                                                            |             |
| 3 – Warm Reset                | A restart of the system is performed by software that does not involve powering the system off or activating the entity’s reset circuitry.                                                     |             |
| 4 – Manual Hard Reset         | A restart is initiated by the user activation of a mechanical device (for example, pressing a button) and bypasses runtime software.                                                           |             |
| 5 – Manual Warm Reset         | A restart is initiated by the user activation of a mechanical device (for example, pressing a button) and does not involve powering the entity off or activating the system’s reset circuitry. |             |
| 6 – System Restart            | A restart of the entity is initiated by entity hardware components and accomplished by activating the system’s reset circuitry.                                                                |             |
| 7 – Watchdog Timeout          | A restart of the entity is initiated in response to a detected system hang condition.                                                                                                          |             |

| Set ID 196 Boot Progress                                                   | System firmware or software booting status.                                               | Notes/Usage                          |
| :------------------------                                                  | :-------------                                                                            | :-------                             |
| 0 - Unknown                                                                | Unknonwn, not a defined value                                                             | Initial state before UEFI is entered |
| 1 ... 5 (Not Used)                                                         | Not Used                                                                                  |                                      |
| 6 – User Authentication                                                    | The boot process is processing the user authentication.                                   |                                      |
| 7 – User-Initiated System Setup                                            | System firmware or BIOS has entered the user system firmware or BIOS configuration setup. |                                      |
| 8 ... 20 (Not Used)                                                        | Not Used                                                                                  |                                      |
| 21 – Starting Operating System Boot Process (for example, calling INT 19h) | The system firmware or BIOS is booting the operating system.                              | Marks transition from UEFI to OS     |
| 22 ... 26 (Not Used)                                                       | Not Used                                                                                  |                                      |

Reference:

- <https://www.dmtf.org/dsp/DSP0249> DMTF DSP0249 State Set Specification
- <https://www.dmtf.org/dsp/DSP0248> DMTF DSP0248 PLDM for Platform Monitoring and Control Specification

### PLDM RDE

Reference:

- <https://www.dmtf.org/dsp/DSP0218> DMTF DSP0218 PLDM for Redfish Device Enablement

### SPDM

SPDM runs over MTCP/I2C, PCIe DOE, and potentually MTCP/USB.  Not that some physical layers do not support Asyncronus Event Notifications (AEN).  SPDM is initiated by the platform BMC to the xPU.

The platform BMC can use SPDM to get a device certificate or alias certificate from an xPU and challenge that xPU to verify it has the associated private key.  If the platform BMC supports mutual authentication, the xPU can get a device certificate for the platform BMC and challenge it.  The use case for SPDM is not clear.  It makes sense for the platform BMC to validate the xPU if the platform BMC is considered thw primary Root of Trust (ROT).  If the xPU is an independent or primary ROT it may not make sense to have the platform BMC validate it.  Also note that verification of authenticity requires a trust store of CA certs to be kept on the device doing the verification.

If mutual authentication is supported it may make sense to privilage some operations from platform BMC to XPU for example NMI, reset or graceful shutdown requests.

Using an SPDM encrypted session might be a good way share credentials between the xPU and the platform BMC to be used for protocols that require credentials (i.e. Redfish).

Reference:

- https://www.dmtf.org/dsp/DSP0274 DMTF DSP0274 Security Protocol and Data Model (SPDM) Specification

### I2C

I2C on it's own is not useful unless there is a protocol defined.  

### NC-SI OEM

NC-SI 1.2 draft has no standard equivilant of the PLDM State Set Specification so OEM specific extensions would be required.

### usb

Could use PLDM State Sensor over PLDM over MTCP over USB
Reference:

- <https://www.dmtf.org/sites/default/files/standards/documents/DSP0283_0.1.5WIP10.pdf> (WIP) DMTF MCTP over USB Binding Specification

### others

## Out-of-band xPU BMC
