# Bootup coordination approaches

Reference [BOOTSEQ.md](./BOOTSEQ.md)

## Use cases
This proposal applies to the following use cases:
1.	One or more xPUs in a single server
2.	A single xPU attached to multiple CPUs within the same server
3.	The xPU is powered on at the same time as the server BMC using Auxiliary power
4.	The server BMC controls when the xPU and server CPUs are powered with Primary power
5.	The xPU boots from the network or on-board non-volatile storage
6.	xPUs with and without a local Add-in Management Controller (AMC)

The multi-host case where one or more xPUs are attached to multiple servers within an enclosure is deferred to another exercise.  This requires a Chassis Management Controller (CMC).

## Assumptions
1.	While the xPU is powered by 3.3Vaux only
  a.	OOB communications from the server BMC shall be limited to accessing the FRU (EEPROM) and/or security chip on the xPU.
  b.	There shall be provide no IB communication from the server CPU
2.	The xPU may ignore transitions in the PERST# signal
3.	The xPU may require separate mechanisms to reset its PCIe interface and compute complex
4.	The xPU may start to boot as soon as it detects the stable application of +12V and +3.3V
5.	The xPU is agnostic as to if +12V and +3V are derived from Primary or Auxiliary power
6.	The xPU has no dependencies on the status of the server OS
7.	The xPU shall be fully booted before UEFI/BIOS hands execution over for the server OS to start booting
8.	The xPU may or may not be the only network data path (tenant) to the host OS
9.	Downstream Port Containment (DPC) is not a practical mechanism to hot add PCIe functions during POST
10.	When an xPU is equipped with an AMC it shall be subservient to the server BMC

## In-band PCIe

- what problems we have today with PCIe, timeouts, errors, retries,...

## Out-band via BMC

- i2c
- ncsi
- usb
- others...
