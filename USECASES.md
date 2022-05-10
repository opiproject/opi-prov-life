# Deployment scenarios

- Private network/cloud – (DHCP/DNS trusted)
               Telco
               Enterprise
- Colo – DHCP not trusted
- Other multitenant

# Usage scenarios

- Same org owns host and DPU
- Different org owns host and DPU, or uses host and DPU
- What are the needs/usages of third party OS and Apps on DPU?
- Pre-installed OS works in some scenarios, bare metal in others

# SKUs
- BMC vs no BMC on DPU; need to cover both scenarios
- Availability of ports – some will have management port, others not

= = = =

# Usages

- Distributed firewall
- Virtual switch offloading
- Other networking services
- Storage offloads for composable systems
- NVMEoF
- …

# Derived requirements

- ZTP
               Agent running on DPU/CPU cores
               Installed by default – what are the mandatory default agents/apps/OS 
               sZTP – when used?
               FDO – when used? 
- Telemetry for service deployment
- Secure wipe? (reprovisioning) 
               What does it erase
                              What memory devices
               Leave UEFI
                              What is present on a “bare metal” DPU

- Authenticating DHCP –
               List of MAC addresses (but could be spoofed)
               Perhaps better to use PKI to authenticate provision server specified in options
- Pre-provisioning step for provisioning server authentication
               Could be done at factory
               Reseller
               Integrator
               Telco (could be end owner)
- One method: first touch claims ownership (default credentials?)

# Notes
- Sonic community (there is precedent) – ZTP – the initial password is provisioned rather than default
- 802.1AR – device identity reference
- ONIE (precedent for provisioning/default OS) – small OS(linux) with ZTP client
- Other precedents – clients define installation
               PXE
               Boot from host (noted as less preferred)

