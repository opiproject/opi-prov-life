# Notes from November 7th, 2023

## Summary

- Question was raised about recording the meeting in Gal's absence. Consensus was that we needed to have an interim chair named, which would be brought up at the next TSC meeting.

- Bracha provided a brief update on the PLDM states and next steps

- Further PLDM state discussion ensued, and we believe (as a consensus) that we need to establish a new register by working with the DMTF to standardized as such. The proposed state table from the combined states covered in pldm.md (PLDM column)

- Question was raised should reboots of xPUs be thought of/modeled as a hotplug event.

- One idea raised was should the DPU need to reboot, PCI Express Advanced Error Reporting could be utilized.

- Toshi indicated he would take a look at PCIe AER in the Linux kernel to see if it could be a viable option, if time permits, and how that might interact with libvirt (in terms of PCI pass-through).
