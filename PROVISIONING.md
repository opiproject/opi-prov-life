# Discovery and Provisioning

## Docs
- please look at https://github.com/silicom-ltd/uBMC
- please look at https://opencomputeproject.github.io/onie/overview/
- please look at https://fidoalliance.org/intro-to-fido-device-onboard/
- please look at https://www.rfc-editor.org/rfc/pdfrfc/rfc8572.txt.pdf (Secure Zero Touch Provisioning (SZTP))
- please look at https://pypi.org/project/sztpd/
- please look at https://watsen.net/docs/sztpd/0.0.11/admin-guide/

## Re-Provisioning

tbd... re-celling, faults...

## Manual Provisioning

tbd...

## USB/Virtual media Provisioning

- Provisioning server contacts xPUs BMC (i.e. via redfish)
  - Question: how can we get list of IPs / MACs and credentials ? Manual ?
  - Question: can we also do the DHCP discovery of the BMC and initiate the provisioning from the BMC itself ?
- Provisioning server changes boot order of the xPU
- Provisioning server maps Virtual Media with ISO image to provision
- Provisioning server causes reboot after or xPU reboots itself

## Automatic Provisioning

TODO: need flow chart here
Question: how to call it PXE ? iPXE ? HTTPs?

- Device is powered on
  - Question: how? BMC ? Always on?
- Device runs DHCP client to send a DHCP request packet to the DHCP server.
  - Question: DHCP client runs from Linux or from UEFI ?
  - Question: Is DHCP can be assumed ? What else mDNS? Static IP ? Mac ?
  - Question: Is DHCP relay required ? in case we are on the different network ?
-  DHCP server assigns an IP address, a default gateway, and the IP address or domain name of the "bootstrap server" to the device.
  - Question: is "bootstrap server" in the cloud ? local in datacenter ? remote ? vm/container ?
- Device performs two-way authentication with the bootstrap server and establishes an HTTPS connection with the bootstrap server
  - Question: preconfigured certificates? how they distributed? what is alternative to certificates?
  - Question: one way or mutual (two-way) authentication is required ?
- Device downloads the system software, configuration file, and so on...
  - Question: bootstrap server can point to deployment file server? or they are the same?
- Device reboots into newly installed software
  - Question: if the version is the same, can the entire process skip ? where this happens?

## Progress / Monitoring

- Inventory Query or Broadcast... 
 - So the provisioning server can present the xPUs and their limited
 - Which DPUs in Which Host
 - Virtual location in the host ()
 - Credentials
 - Vendor, SN/PN, ...
 - What protocols available
 - Capabilities 

- Monitoring/Status of provisioning itelf...
  - like bad certificates, other problems, ...

## What OPI produces?

- OPI can produce secure provisioning server/vm/container that follows some well defined SPEC
- Any other implementation of the secure provisioning server/vm/container is also acceptable if follows same SPEC
- xPU vendors will adopt their implementation to meet this secure provisioning server/vm/container
- Provisioning companies/customers will use API defined to the secure provisioning server/vm/container to integrate in their existing provisioning methods

## TBD

tbd