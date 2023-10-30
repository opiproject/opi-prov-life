# PLDM State sensors

This page is work in progress, aggregating all the information we have so far with all the proposals.
We are iterating on this with pull requests until reaching final state.
Once finilized, we are taking this proposal to DMTF PLDM cometee.

## Terms

from [here](./README.md#terms)

## Diagram

from [here](./COORDINATION.md#diagram)

```mermaid
sequenceDiagram
    autonumber
    participant ServerBmc
    participant ServerBios
    box rgb(100,100,100) xPU
       participant xPUmngmnt
       participant xPUCpu
    end
    ServerBmc->>ServerBios: Power On
    ServerBmc->>xPUmngmnt: Power On
    xPUmngmnt->>xPUCpu: Power On
    ServerBmc->>xPUmngmnt: FRU request
    xPUmngmnt->>ServerBmc: FRU response
    ServerBmc->>ServerBios: hold enumeration
    xPUCpu->>xPUmngmnt: Update Boot status 1
    xPUCpu->>xPUmngmnt: Update Boot status 2
    xPUCpu->>xPUmngmnt: Update Boot status 3
    ServerBmc->>xPUmngmnt: Get Boot status
    xPUmngmnt->>ServerBmc: xPU ready
    ServerBmc->>ServerBios: continue enumeration
```

## Reference

- <https://www.dmtf.org/dsp/DSP0249> DMTF DSP0249 State Set Specification
- <https://www.dmtf.org/dsp/DSP0248> DMTF DSP0248 PLDM for Platform Monitoring and Control Specification

## Exisitng PLDM State sensors

from [here](./COORDINATION.md#pldm-state-sensors---pldm)

Useful Entity IDs:

- 31 System Firmware (for example, BIOS/EFI)
- 32 Operating System
- 166 PCI Express Bus

Useful State definitions:

- 1 Health State
- 2 Availability
- 129 Software Termination Status
- 192 Boot/Restart Cause
- 196 Boot Progress

### 1 Health State

| Set ID 1 Health State  | Represents the current health of the entity.                                                                           | Notes/Usage |
| :-----                 | :-----                                                                                                                 | :-----      |
| 1 – Normal             | The entity is at a normal state of health.                                                                             |             |
| 2 – Non-Critical       | The entity is not at a normal state of health, but is still operational.                                               |             |
| 3 – Critical           | The entity is at a critical state of health. The entity may have suffered permanent damage, and may not be functional. |             |
| 4 – Fatal              | The entity is at a fatal state of health.                                                                              |             |
| 5 ... 10 (not used)    | Not Used                                                                                                               | These states are for thermal |

### 2 Availability

| Set ID 2 Availability | The operational state of the entity.      | Notes/Usage |
| :-----                | :-----                                    | :-----      |
| 1 – Enabled           | The entity is in an enabled state.        |             |
| 2 – Disabled          | The entity is in a disabled state.        |             |
| 3 – Shutdown          | The entity has been shut down.            |             |
| 4 – Offline           | The entity is in an offline test.         |             |
| 5 – In Test           | The entity is in a test mode.             |             |
| 6 – Deferred          | The entity has been deferred to function. |             |
| 7 – Quiescent         | The entity is quiescent to function.      |             |
| 8 – Rebooting         | The entity is currently rebooting.        |             |
| 9 – Resetting         | The entity is resetting.                  |             |
| 10 – Failed           | The entity is in a failed state.          |             |
| 11 – Not Installed    | The entity is not installed.              |             |
| 12 – Power Save Mode  | The entity is in a power save mode.       |             |
| 13 – Paused           | The entity is paused.                     |             |
| 14 – Shutting Down    | The entity is shutting down.              |             |
| 15 – Starting         | The entity is starting or initializing.   |             |
| 16 – Not Responding   | The entity has stopped responding.        |             |

### 129 Software Termination Status

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

### 192 Boot/Restart Cause

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

### 196 Boot Progress

| Set ID 196 Boot Progress                                                   | System firmware or software booting status.                                               | Notes/Usage                          |
| :------------------------                                                  | :-------------                                                                            | :-------                             |
| 0 - Unknown                                                                | Unknonwn, not a defined value                                                             | Initial state before UEFI is entered |
| 1 ... 5 (Not Used)                                                         | Not Used                                                                                  |                                      |
| 6 – User Authentication                                                    | The boot process is processing the user authentication.                                   |                                      |
| 7 – User-Initiated System Setup                                            | System firmware or BIOS has entered the user system firmware or BIOS configuration setup. |                                      |
| 8 ... 20 (Not Used)                                                        | Not Used                                                                                  |                                      |
| 21 – Starting Operating System Boot Process (for example, calling INT 19h) | The system firmware or BIOS is booting the operating system.                              | Marks transition from UEFI to OS     |
| 22 ... 26 (Not Used)                                                       | Not Used                                                                                  |                                      |

## Dell NC-SI implementation

from [here](../architecture/Dell%20NC-SI%20OEM%20Commands%20for%20smartNICs.pdf)

The following table defines the SN State values
| Value | Name | Description |
| :-----| :----| :-----------|
| 0 | Reset | CPU is in reset / Boot ROM
| 1 | Firmware #1 | CPU has passed FW checkpoint 1
| 2 | Firmware #2 | CPU has passed FW checkpoint 2
| 3 | UEFI | CPU has entered UEFI
| 4 | OS Booting | CPU has entered OS
| 5 | OS Running | OS is running
| 6 | OS Halted/Shutdown | OS is halted or shutdown
| 7 | Updating | Update in Progress
| 8 | OS Crash Progressing | OS Crash Dump in progress
| 9 | OS Crash Complete | OS Crash Dump complete
| Other | Reserved | Reserved

## Nvidia proposal to extend 196 Boot Progress

from [here](https://opi-project.slack.com/archives/C0342L6T7EC/p1693938501126579)

| Value/Name                                                   | Description                                               | Notes/Usage                          |
| :------------------------                                                  | :-------------                                                                            | :-------                             |
| 0 - Reset/Boot-ROM                   | The device has just been powered on or reset, and it's initializing basic hardware and loading the first firmware mutable FW components. |
| 1 - Boot stage 1                     | FMC (First Mutable Code) is running. |
| 2 - Boot stage 2                     | The device has progressed further in the boot process, executing additional instructions to load pre-OS SW. |
| 3 - UEFI                             | The device is transitioning into the Unified Extensible Firmware Interface (UEFI) environment. |
| 4 - OS starting                      | The operating system (OS) is being loaded and initialized on the device. |
| 5 - OS is running                    | The operating system has successfully started and is now running normal operations. |
| 6 - Low-Power standby                | The device has entered a low-power standby mode or is placed into idle state, conserving energy while remaining operational. |
| 7 - Firmware update in progress      | The device's firmware is being updated with new image. |
| 8 - OS Crash Dump in progress        | The operating system is in the process of capturing diagnostic information about an OS crash or error that has occurred. |
| 9 - OS Crash Dump is complete        | The operating system has concluded recording a crash dump. |
| 10 - FW Fault Crash Dump in progress | The device's firmware is collecting diagnostic data related to a FW fault or error in its operation. |
| 11 - FW Fault Crash Dump is complete | The firmware has finished collecting diagnostic data about a fault. |

For each state we want to define a "state result", which basically explains what we can expect after each state is done, I will take it internally and come with a proposal but feel free to do the same.

Each vendor should go over the list, see how it fits into his current flow and if there are new states which are needed or states that are unnecessary.

On top of the states we want to try and create a document in PLDM describing "how should a xPU be managed with PLDM".

## Comparison of PLDM, Dell and Nvidia proposals

| Dell           | Nvidia                          | PLDM                                                                                                                 | Notes                                                  |
| :------------  | :------------                   | :------------                                                                                                        | :------------                                          |
| 0 - Reset      | 0 - Reset/Boot-ROM              | ?.xPUBootPrg = 0;                                                                                                    | Hold Boot. Initial state on reset                      |
| 1 - FW #1      | 1 - Boot stage 1                | ?.xPUBootPrg = 1;                                                                                                    | Hold Boot. Informational                               |
| 2 - FW #2      | 2 - Boot stage 2                | ?.xPUBootPrg = 2;                                                                                                    | Hold Boot. Informational                               |
| 3 - UEFI       | 3 - UEFI                        | ?.xPUBootPrg = 3; EFI(21).Avail(2) = Enabled(1)                                                                      | Hold Boot. Informational                               |
| 4 - OS Booting | 4 - OS starting                 | ?.xPUBootPrg = 4; ?.BootPrg(196) = StartingOS(21); OS(32).Avail(2) = Starting(15)                                    | Hold Boot. Informational                               |
| 5 - OS Running | 5 - OS is running               | ?.xPUBootPrg = 5; OS(32).Health(1) = Normal(1); PCIe(166).Avail(2) = Enabled(1)                                      | Release hold.  Hold PCI enumeration until this state   |
| 6 - OS Halted  | 6 - Low-Power standby           | ?.xPUBootPrg = 6; OS(32).Avail(2) = Shutdown(3); OS(32).Term(129) = Shutdown(7)                                      | Graceful shutdown.  Safe to power off                  |
| 7 Updating     | 7 - Firmware update in progress | ?.xPUBootPrg = 7;                                                                                                    | Do not power off.  Informational                       |
| 8 OS Crashing  | 8 - OS Crash Dump in progress   | ?.xPUBootPrg = 8; OS(32).Health(1) = Fatal(4); OS(32).Term(129) = TermDetected(2)                                    | Do not power off.  Informational                       |
| 9 OS Crashed   | 9 - OS Crash Dump comple        | ?.xPUBootPrg = 9; OS(32).Health(1) = Fatal(4); OS(32).Avail(2) = Failed(10); OS(32).Term(129) in {3,4} critical stop | Requesting cold boot on next reset.  Safe to power off |
| N/A            | 10 - FW Fault in progress       | ?.xPUBootPrg = 10;                                                                                                   | Do not power off.  Informational                       |
| N/A            | 11 - FW Fault completete        | ?.xPUBootPrg = 11;                                                                                                   | Requesting cold boot on next reset.  Safe to power off |
