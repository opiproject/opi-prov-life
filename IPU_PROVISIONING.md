# IPU Provisioning

Reference [BOOTSEQ.md](./BOOTSEQ.md)

## TLDR

From a provisioning perspective, an Infrastructure Processing Unit (IPU) works
as a Multi-Host NIC with an embedded core complex as one of the hosts. A
Multi-Host NIC is split between 2 or more host interfaces (commonly over PCIe)
and 1 or more Network Interfaces (commonly over Ethernet).  Host Devices
are divided into '''Data Plane''', '''Control Plane''', and '''Management'''
which can be individually passed to different hosts and have their own
capabilities, reset and security domains, where each device is hardware
isolated from each other.

## Device Types

Provisioning follows the same rules as a Multi-Host NIC, based on a set of
device types:
 * '''Data Plane Devices''': Virtio-net, Virtio-blk, NVMe and idpf are data
plane devices exposed to attached hosts.  Resetting a data plane device will
reset the device only, and will leave all other devices (and all other hosts)
unchanged.
 * '''Control Plane Device''': A control plane device (usually on the embedded
core complex) is capable of resetting any of the host interfaces (e.g., PCIe)
of the Multi-Host NIC. It is also able to request a full device reset from the
management controller.
 * '''Management Device'';': A Chassis Management Controller (CMC) and/or an
Integrated Management Controller (IMC) is given a Management Device capable of
managing and resetting the entire IPU.

## Host Perspectives

From the point of view of an attached host (including the embedded complex):

 * Startup Enumeration: The set of devices available to that Host at startup
will be found right at PCIe enumeration time.
 * Data Plane Reset: Resetting a data plane PCIe device will reset the state
for that specific data plane device, and will not have any effect on other
devices or on any attached host interfaces.
 * Network Data Plane Devices: Virtio-net and idpf can be used to send and
receive network traffic, including stateless offloads such as TSO, RSS, etc.
 * Storage Data Plane Devices: Virtio-blk/scsi and NVMe can be used to do
block storage reads, writes, flushes, etc. on a disk assigned to that host.
 * Host Boot: Attached hosts can boot from network or storage data plane
devices.  Each device's driver verifies that the boot mechanism is ready
before proceeding with the host boot, to ensure that the boot disk accessible
(often times coming from over the network).
 * Control Plane is Optional: One host is responsible for the control of
traffic between hosts and between data plane devices. The remaining hosts use
the data plane devices assigned to them through their PCIe interface. A host
with a Control Plane device can also configure the Ethernet Line Side
including setting of speed, serdes, link layer functions, etc.
 * IPU Reset: Resetting of the IPU is done through the Management Controller,
either by the attached CMC or an integrated IMC. Control Plane devices can
request for the IPU to be reset by doing an in-band request to the
management controller.
 * Multi-Socket: Connecting up multiple PCIe connections to the same core
complex but on different sockets can enable multiple Data Plane devices from
the same IPU to be connected to different sockets.
 * Visibility:  A host with data plane devices (only) has no visibility into
the data plane implementation (could be an IPU, DPU, software backend, etc.).
A host with a control plane device has control of everything in the IPU that
has been granted to it by the management device.  The Management Controller
with access to the management device has full control of the IPU and all of its
host and network interfaces.

## Platform Perspectivies

From the point of view of the server platform:
 * Power: The IPU can be powered from the PCIe goldfingers or be given its own
independent power, depending on the design of the platform.
 * Boot Sequence: When power is shared, booting of the attached host(s) and
IPU will happen in parallel, with host data plane and control plane set up
at PCIe enumeration time. Independent power can enable serialized boot
sequencing which is also supported by the IPU.
 * I2C: On the I2C over the PCIe goldfingers the IPU looks identical to a NIC
w/ FRU and power sensors. No additional information is required over the I2C
to put an IPU into a server.
 * NC-SI: Most IPU cards include an optional cable containing an NC-SI
connector meant to be plugged into a platform BMC.  Not required for operation.
 * Console: The IPU has a console port, often exposed alongside NC-SI or over
a debug connector on the faceplate. Not required for operation.
 * Dynamic Data Plane Devices: BIOS and IPU can be configured to support
dynamic hot plug of data plane devices on the PCIe. Not required for operation.
 * Host Reboot/Failure: Individual hosts can stall, stop, reboot, crash, 
lose power, etc., and this will have no direct effect on the operation of the
IPU or any of the other attached hosts if their power continues to be supplied.
 Loss of the host with the control plane device likely will have an effect on
the functionality of data plane devices on other hosts. 
 * Platform BMC: The platform BMC may connect to the IPU over NC-SI or any of
the other IPU connections such a Gigabit Ethernet or debug connector.
Connection to the Platform BMC is not required for IPU operation.
 * Chassis Management Controller (CMC): The IPU Management Device controls the
reset of the IPU, as well as the allocation of resources to attached hosts.
Advanced operations beyond the Multi-Host NIC default configuration requires
connectivity to the chassis and/or integrated management controller.
 * IPU Reboot/Failure: IPU failure is visible through control plane and
management devices.  Hosts with only data plane device visibility do not have
visibility into an IPU reboot or failure, beyond their data plane devices
losing traffic or stalling new requests. An IPU can be rebooted/reinitialized
and/or have its software updated while keeping the data plane devices on
attached hosts up and running without prolonged service interruption. Loss
of the IPU and/or loss of the host handling the control plane device will
likely cause the data plane devices on other hosts to stop responding.

# Vendor Survey

 - What is the role of a BMC on the xPU regarding startup?
   - The platform BMC is not required for IPU startup or operation.
 - What casues the ARM cores to start running? Application of 12V? Release of PCIe RESET? Write to a location via NC-SI or I2C? ??
   - Embedded cores (either x86 or ARM based) are their own host complex and
are controlled via the Management Controller (CMC, IMC, etc.)
 - How does PCIe reset affect the xPU's PCIe interface and the ARM cores?
   - PCIe reset of a data plane device only resets that specific device
   - Embedded cores can reset themselves are are not resettable by any other
host via PCIe alone.
 - What happens to the xPU's PCIe interface when the ARM cores lockup?
   - Loss of the embedded core complex has no direct impact on other host
complexes attached to the IPU. Loss of the control device likely will cause
data plane devices on other hosts to stop responding.
 - What is the maximum time from release of PCIe RESET until the xPU responds to config cycles?
   - Data plane devices will start responding once their driver has determined
that the backend is ready via a Driver Ready Check.
   - The time it takes to initialize the networking and storage pathways set up
by the IPU is a function of the programming running on the host containing the
control plane device.
 - How does the xPU determine if it should boot to its OS or some diagnostic FW?
   - This is controlled by the Management Controller.
 - Is there a difference in xPU PCIe behavior between an OS running or diagnoostic FW?
   - The Management Controller and the host containing the Control Plane device
are independent, separate entities.
 - What is the method for determining the xPU boot status via PCIe, I2C and NC-SI?
   - Data plane devices will show ready when the backend clears the Device
Ready check. Each host can boot/reboot/reset/etc. independent of one another.
The Management Controller is ultimately the arbiter of shared resources and
controls the full IPU reset.




