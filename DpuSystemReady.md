# xPU Base FW and Security requirements

## TLDR

This page lists xPU base firmware and security requirements and recommendations. Having standard firmware and security interfaces ensures that standard off-the-shelf operating systems and hypervisors have the minimal interfaces necessary to run on the xPUs in a consistent manner. Since some of these requirements may be CPU architecture specific, xPUs based on different CPU ISAs may have different requirements.

## Terms

| Term | Description
|------|----------------------------------------------------------------------|
| ACS  | Arm Architecture Compliance test Suite                               |
| BBR  | Arm Base Boot Requirements specification                             |
| BSA  | Arm Base System Architecture specification                           |
| BBSR | Arm Base Boot Security Requirements specification                    |
| UEFI | Unified Extensible Firmware Interface previously referred to as BIOS |
| xPU  | SmartIO, SmartNIC, DPU, or IPU                                       |

## _**Proposal - To be discussed**_:

1. xPUs with Arm-based processors are recommended to achieve [Arm SystemReady ES v 1.4](https://www.arm.com/architecture/system-architectures/systemready-certification-program/es)(or newer) certification .
2. xPUs with Arm-based processors are required to be compliant with the following specifications. Requirements are to be verified using [SystemReady ES ACS v1.2.0](https://github.com/ARM-software/arm-systemready/tree/main/ES) (or newer) 
     - BSA ver 1.0c (or newer).
     - SBBR recipe in BBR v1.0 (or newer), which includes implementing UEFI, ACPI, and SMBIOS.
     - BBSR v 1.2 (or newer).

## Arm SystemReady certification program

[Arm SystemReady](https://www.arm.com/architecture/system-architectures/systemready-certification-program) is a compliance certification program based on a set of hardware and firmware standards: [Base System Architecture (BSA)](https://developer.arm.com/documentation/den0094/latest) and [Base Boot Requirements (BBR)](https://developer.arm.com/documentation/den0044/latest), as well as an [Architecture Compliance Suite (ACS)](https://github.com/ARM-software/arm-systemready). This program tests and certifies that systems meet the SystemReady standards, giving confidence that operating systems (OS) and subsequent layers of software just work. 

The SystemReady certification and testing requirements are specified in the [Arm SystemReady Requirements Specification (SRS)](https://developer.arm.com/documentation/den0109/latest).


## Arm SystemReady Security Interface Extension (SIE)
The Arm SystemReady Security Interface Extension (SIE) provides a way to certify that UEFI secure boot, secure firmware update, and TPM interfaces are implemented, as prescribed by the [Arm Base Boot Security Requirements (BBSR)](https://developer.arm.com/documentation/den0107/latest/) specification. An [Architectural Compliance Suite (ACS)](https://github.com/ARM-software/arm-systemready) is available to verify the compliance of a firmware implementation to BBSR.


     
