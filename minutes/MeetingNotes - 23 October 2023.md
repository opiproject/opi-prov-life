# Notes from October 23rd, 2023

## Summary

Consensus was reached that: In order to make forward progress, we believe we need to think/ponder and discuss along the following lines:

- Do we need to view multiple devices as a first class model to be supported in a single host? i.e. Do we need to consider and handle High Availability concerns? Specifically if we need to go down the PLDM route to coordinate halt/reboot/power-off coordination with the device, as we may encounter some constraints like with timing.

- Overall, are there any concerns with the use of "need to reboot" signaling taking place with PLDM as well? Specifically we need the hardware vendors to discuss internally and inform the group as guidance as the next step is likely the creation of flow diagrams describing the flow given the communication modeling.

- Would use of PLDM be a limiting factor with future transport modeling? The concern of future USB transport support for device communications was raised, but we don't know if it would, or would not be. Also likely an internal topic for hardware vendors.

