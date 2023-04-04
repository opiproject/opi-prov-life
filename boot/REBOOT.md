# DPU and Host reboot coordination approaches

Reference [BOOTSEQ.md](../BOOTSEQ.md)

## TLDR

This page is taking assumptions and use cases from [BOOTSEQ.md](../BOOTSEQ.md) page and tries to deep dive on possible solutions of the only DPU and Host reboots parts.

Initial bootup part covered [COORDINATION.md](./COORDINATION.md): Server is Powered On, DPU receives power and starts booting, Host OS should wait for DPU to finish booting.

## Terms

see <https://github.com/opiproject/opi-prov-life/blob/main/BOOTSEQ.md#terms>

## 1: DPU reboots

### Why

* Host can yank DPU power if Host is not aware that DPU is rebooting
* Host can use the notification to migrate workloads to another host or to another DPU in the same host, while DPU is rebooting
* Host can apply some kind of "protection" to prevent DPC crushing the Host OS (surprise removal doesn't work well in NICs)
* Host can allow DPU to collect core dump
* Host can stay alive or gracefully reboot once DPU is rebooted

### What

* DPU should notify to Host that is is about to reboot
* In a trusted environment Host can ask for extension
* DPU will reboot itself anyways after timeout/wd expires

### How

* Out-band via platform BMC: maybe using techiniques from [BOOTSEQ.md](../BOOTSEQ.md)
* In-Band TBD: maybe PCIe, IRQ, ...

## 2: Host reboots

* Host can choose to reboot in different ways (shell, ipmi)
* Maybe we can use somehow PCIe PERST signal to notify DPU that host reboots
* DPU would like to free some resources/context assoiciated with that Host
* DPU can re-provision it for next Tenant on reboot
* DPU can deny/prevent from Host reboot
* How DPU reaction changes if it is installed in a slot providing persistent power, implying it should be independent of host?
* TBD
