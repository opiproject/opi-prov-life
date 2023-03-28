# Survey 2/17/2023 responses

**1. What is the role of a BMC on the xPU regarding startup?**
BMC is powered by auxiliary power, which is always on, and controls the reset signal to the AP cores. BMC has access to the management port, SPI, eMMC, etc.
**2.  What causes the ARM cores to start running? Application of 12V? Release of PCIe RESET? Write to a location via NC-SI or I2C? ??**
ARM cores are managed by internal coprocessors. BMC can communicate with them either via I2C or I3C and controls the reset signal. PCI reset does not drive their operation.
**3. How does PCIe reset affect the xPU's PCIe interface and the ARM cores?**
The domains are isolated. Only the PCI interface will reset and not the ARM cores. However, software should be made aware that reset is occurring to avoid invalid transactions during the reset process.
**4. What happens to the xPU's PCIe interface when the ARM cores lockup?**
Nothing. The PCIe interface will still be alive
**5. What is the maximum time from release of PCIe RESET until the xPU responds to config cycles?**
We are PCI spec compliant and complete the training and link up within the spec
**6. How does the xPU determine if it should boot to its OS or some diagnostic FW?**
Our current implementation is software controlled. However, the platform BMC can control the behavior as required
**7. Is there a difference in xPU PCIe behavior between an OS running or diagnostic FW?**
No, the PCI link up behavior is the same in both modes
**8. What is the method for determining the xPU boot status via PCIe, I2C and NC-SI?**
We support all three transport mechanisms, and can support various protocols to communicate the boot checkpoint status to the host
