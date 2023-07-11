## Remote Host/BMC Inventory

Host host/platform BMC access inventory information from the DPU/IPU ?

- NC-SI
- PLDM
- VPD
- other...

VPD Example from Nvidia:

```bash
$ lspci -s ca:00.0 -vvv | grep -A 11 "Vital Product Data"
        Capabilities: [48] Vital Product Data
                Product Name: BlueField-2 E-Series SmartNIC 100GbE/EDR VPI Dual-Port QSFP56, PCIe Gen4 x16, Crypto Enabled, 16GB on-board DDR, FHHL
                Read-only fields:
                        [PN] Part number: MBF2M516A-EEEOT
                        [EC] Engineering changes: A1
                        [V2] Vendor specific: MBF2M516A-EEEOT
                        [SN] Serial number: MT2022X19080
                        [V3] Vendor specific: 805647c144a9ea1180000c42a198b662
                        [VA] Vendor specific: MLX:MN=MLNX:CSKU=V2:UUID=V3:PCI=V0:MODL=BF2M526A
                        [V0] Vendor specific: PCIeGen4 x16
                        [RV] Reserved: checksum good, 1 byte(s) reserved
                End
```
