# Discovery and Provisioning

## Docs
- please look at https://github.com/silicom-ltd/uBMC
- please look at https://opencomputeproject.github.io/onie/overview/
- please look at https://fidoalliance.org/intro-to-fido-device-onboard/
- please look at https://www.rfc-editor.org/rfc/pdfrfc/rfc8572.txt.pdf (Secure Zero Touch Provisioning (SZTP))
- please look at https://www.rfc-editor.org/rfc/pdfrfc/rfc8366.txt.pdf (A Voucher Artifact for Bootstrapping Protocols)
- please look at https://pypi.org/project/sztpd/
- please look at https://watsen.net/docs/sztpd/0.0.11/admin-guide/
- please look at https://nvlpubs.nist.gov/nistpubs/CSWP/NIST.CSWP.09082020-draft.pdf (NIST)
- please look at https://www.rfc-editor.org/rfc/rfc8995.html (BRSKI)
- please look at https://access.redhat.com/articles/6804281#install-assisted-installer-on-the-installer-node-using-podman-5
- please look at https://github.com/jparrill/ztp-the-hard-way (RHEL)
- please look at https://access.redhat.com/documentation/en-us/openshift_container_platform/4.9/html/scalability_and_performance/ztp-deploying-disconnected

## Re-Provisioning

tbd... re-celling, faults...

## Manual Provisioning

tbd...

## USB/Virtual media Provisioning

Use case: small scale, unique, specialized deployments ?

- Provisioning server contacts xPUs BMC (i.e. via redfish)
  - Question: how can we get list of IPs / MACs and credentials ? Manual ?
  - Question: can we also do the DHCP discovery of the BMC and initiate the provisioning from the BMC itself ?
- Provisioning server changes boot order of the xPU
- Provisioning server maps Virtual Media with ISO image to provision
- Provisioning server causes reboot after or xPU reboots itself

## Automatic Provisioning

TODO: need flow chart here
Question: how to call it PXE ? iPXE ? HTTPs?
Use case: large scale deployments (where automation and security are major drivers) ?

- Device is powered on
  - Question: how? BMC ? Always on?
- Device runs DHCP client to send a DHCP request packet to the DHCP server.
  - Question: DHCP client runs from Linux or from UEFI ?
  - Question: Is DHCP can be assumed ? What else mDNS? Static IP ? Mac ?
  - Question: Is DHCP relay required ? in case we are on the different network ?
- DHCP server assigns an IP address, a default gateway, and the IP address or domain name of the "bootstrap server" to the device.
  - Question: is "bootstrap server" in the cloud ? local in datacenter ? remote ? vm/container ?
  - Comment: probably need to support multiple preferences by operators. some facilities have to be actually local, other can be proxied to a centralized location, others yet can be completely centralized without a local presence.
- Device contacts the bootstrap server to get the certificate, bootstrap server will facilitate the request towards CA
  - Take a look at SCEP (Simple Certificate Enrollment Protocol) 
  - The communication with bootstrap server doesn't have to be secure at this point
  - Question: preconfigured certificates? how they distributed? what is alternative to certificates? PKI based crypto ?
- If we want to protect stolen/mistaken shiipment (device needs to authenticate network) we have to use Vouchers
  - Take a look at RFC 8366 - A Voucher Artifact for Bootstrapping Protocols
  - More info is here https://github.com/opiproject/opi-prov-life/blob/main/architecture/Zero-Touch-Provisioning%E2%80%94Approaches-to-Network-Layer-Onboarding.pdf
- Device can now establish an HTTPS connection with the bootstrap server using certificates from above
  - Question: one way or mutual (two-way) authentication is required ?
  - Question: three-way trust established here, between device identity, manufacturer, and operator ???
- Bootstrap server can/should point to deployment file server
  - for operational and scaling purposes they should probably be separate. 
  - There are some scaling requirements on the deployment file server where the implementation details can really drive the pattern of redirection between boot and deployment file servers (for instance, you can avoid having to deploy a load balancer in front of deployment file servers by the boot server spreading the load over multiple DNS names).
- Device performs one-two-three-way authentication with the deployment file server and establishes an HTTPS connection with the deployment file server
  - Question: same as above
- Device downloads the system software, configuration file, and so on...
  - Question: bootstrap server can point to deployment file server? or they are the same?
- Device reboots into newly installed software
  - Question: if the version is the same, can the entire process skip ? where this happens?

## Progress / Monitoring

- Inventory Query or Broadcast
  - see [Inventory](INVENTORY.md)

- Monitoring/Status of provisioning itelf...
  - like bad certificates, other problems, ...

Question: how this is implemented and integrated with existing provisioning services?
Question: OPI can produce an agent (container) that runs on DPU for example and collects all this information via redfish, ipmi, lspci, and other specialized tools... And then exposes single common endpoint API so everybody can query it... like MAAS, JESP, RHEL SAAS and so on...

## What OPI produces?

- OPI can produce secure provisioning server/vm/container that follows some well defined SPEC
- Any other implementation of the secure provisioning server/vm/container is also acceptable if follows same SPEC
- xPU vendors will adopt their implementation to meet this secure provisioning server/vm/container
- Provisioning companies/customers will use API defined to the secure provisioning server/vm/container to integrate in their existing provisioning methods

- OPI can also produce an agent (container/service) for Standard Inventory Query that everybody (existing provisioning systems) can query

## TBD

tbd
