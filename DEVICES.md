# Device Management - Note: **Experimental** for Discussion Purposes

## TLDR

Devices attached to the host to/from the infrastructure need to be added,
deleted, and connected up on the infrastructure side
(to the dataplane, hypervisor, etc.).  For supporting VMs & containers these
devices also need to be connected up on the host side.

What used to be a simple virtual 'wire' when everything was on the host, now
requires additional coordination in the infrastructure and on the host.

## Idenfication using PCI

Given that the infrastructure connects to the host through PCI,
the PCI bus is used to uniquely identify a device.  

Devices can be created in 4 different ways:

1. **PCI Enumeration**: At enumeration time devices can exist and be recognized
by the host.
2. **PCI Hotplug**: If hotplug is configured in the host and infrastructure,
devices can be added / removed dynamically using the PCI hotplug mechanism.
3. **SR-IOV**: Physical functions on the PCI bus can expose SR-IOV virtual
functions which can be treated as additional PCI devices by the host or by a
guest virtual machine.
4. **Sub-functions**: Sub-functions logically split a PCI device into
multiple independent functions that can be treated as separate devices.
these sub-functions are not exposed onto the PCI bus.

Given these methods of device creation, a 3-tuple is used to
identify them on the infrastructure side:

* {PCI Physical Function Handle, Virtual Function Index, Sub-function Index}

From the perspective of the infrastructure device, it may not be aware of the
B:D.F (Bus:Device.Function) assigned to it on the host side, which is why a
PF handle is used.  
A virtual function index of zero represents a PCI Physical Function.
A sub-function index of zero represents a device without sub-functions.

This 3-tuple is sufficient to support the first 2 cases where the infrastructure
is serving devices to the host directly through PCI, and there is no explicit
coordination between the infrastructure and the host other than devices being
present on the PCI bus.

## Host/Guest Agent Functionality

In scenarios where there is coordiination between the host and the
infrastructure (for example when VMs and/or containers need to be associated
with their specific devices) an agent needs to be run in each of the hosts and
guests looking to coordinate devices with the infrastructure.

Ultimately these agents are taking their direction from orchestration (either
directly or in-directly via the infrastructure).  A secure communication
channel is established between the agent(s) and the infrastructure where the
following associations are made:

* Mapping from Bus:Device.Function, Sub-function -> Device 3-tuple

Having this mapping on the agent side enables the connectivity of VMs and
containers to the correct device, identified by their 3-tuple.
