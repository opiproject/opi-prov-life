# Lifecycle Management

High level requirements on a DPU/IPU or subsequently a "device"

## Terms

see <https://github.com/opiproject/opi-prov-life/blob/main/boot/README.md#terms>

## Reboot device

example if we using redfish:

```bash
# this is mostly useful after FW update
$ curl -s -k -u <bmc-user>:<password> -X POST -H "Content-Type: application/json" -d '{"ResetType": "ForceRestart"}' https://<bmc-ip-address>/redfish/v1/Managers/<ID>/Actions/Manager.Reset
and
$ curl -s -k -u <bmc-user>:<password> -X POST -H "Content-Type: application/json" -d '{"ResetType": "PowerCycle"}' https://<bmc-ip-address>/redfish/v1/Systems/<ID>/Actions/ComputerSystem.Reset

# see http://redfish.dmtf.org/schemas/v1/Resource.json#/definitions/ResetType
```

## Control device boot source and order

example if we using redfish:

```bash
$ curl -s -k -u <bmc-user>:<password> -X PATCH -H "Content-Type: application/json" -d '{"Boot": {"BootSourceOverrideTarget":"Pxe"} }' https://<bmc-ip-address>/redfish/v1/Systems/<ID>

# see http://redfish.dmtf.org/schemas/v1/ComputerSystem.json#/definitions/BootSource
```

## Update device

* firmware update
* OS update
* software/application update
* NIC BMC (aka AMC) update if exists

example if we using redfish:

```bash
$ curl -k \
     -u <bmc-user>:<password> \
     -X POST \
     -H "Content-Type: application/json" \
     -d '{"ImageURI": "http://someip:someport/path/to/image.bin", "TransferProtocol":"HTTP", "Targets": ["/redfish/v1/UpdateService/FirmwareInventory/BMC<#>"]}' \
     https://<bmc-ip-address>/redfish/v1/UpdateService/Actions/UpdateService.SimpleUpdate

# TODO: post example using MultipartHttpPush...
```

## Recover device

* reset to a known good state, e.g., factory reset
* via network - required
* via host - optional

example if we using redfish:

```bash
$ curl -s -k -u <bmc-user>:<password> -X POST -H "Content-Type: application/json" -d '{"ResetType": "ResetAll"}' https://<bmc-ip-address>/redfish/v1/Managers/<ID>/Actions/Manager.ResetToDefaults
# TODO: add more examples
```

## Set date and time

example if we using redfish:

```bash
$ curl -s -k -u <bmc-user>:<password> -X PATCH -H "Content-Type: application/json" -d '{"DateTime": "2019-04-25T26:24:46+00:00"}' https://<bmc-ip-address>/redfish/v1/Managers/<ID>
{
  "DateTime": "2019-04-25T26:24:46+00:00"
}
```

## Secure boot

example if we using redfish:

```bash
$ curl -s -k -u <bmc-user>:<password> -X POST -H "Content-Type: application/json" -d '{"ResetKeysType": "DeleteAllKeys"}' https://<bmc-ip-address>/redfish/v1/Systems/<ID>/SecureBoot/Actions/SecureBoot.ResetKeys

# see http://redfish.dmtf.org/schemas/v1/Resource.json#/definitions/ResetKeysType

$ curl -s -k -u <bmc-user>:<password> -X PATCH -H "Content-Type: application/json" -d '{"SecureBootEnable":true}' https://<bmc-ip-address>/redfish/v1/Systems/<ID>/SecureBoot

# TODO: add more examples
```

## Account management

example if we using redfish:

```bash
# change default password for the first time
$ curl -s -k -u <bmc-user>:<password> -X PATCH https://<bmc-ip-address>/redfish/v1/AccountService/Accounts/root -d '{"Password": "mynew@password"}'

# list
$ curl -s -k -u <bmc-user>:<password> -X GET https://<bmc-ip-address>/redfish/v1/AccountService/Accounts

# TODO: add more examples, do we also need to add/remove accounts ?
```

## Network management

example if we using redfish:

```bash
$ curl -s -k -u <bmc-user>:<password> -X GET https://<bmc-ip-address>/redfish/v1/Managers/<ID>/Ethernetnterfaces/eth0

$ curl -s -k -u <bmc-user>:<password> -X PATCH -H "Content-Type: application/json" -d '{"InterfaceEnabled": true}'  https://<bmc-ip-address>/redfish/v1/Managers/<ID>/Ethernetnterfaces/eth0

$ curl -s -k -u <bmc-user>:<password> -X PATCH -H "Content-Type: application/json" -d '{"IPv6StaticAddresses": [{"Address": "fe80::966d:aeff:fe76:6e8f", "PrefixLength": 64}]}' https://<bmc-ip-address>/redfish/v1/Managers/<ID>/Ethernetnterfaces/eth0

$ curl -s -k -u <bmc-user>:<password> -X PATCH -d '{"DHCPv4": {"DHCPEnabled": false}}' https://<bmc-ip-address>/redfish/v1/Managers/<ID>/Ethernetnterfaces/eth0

# TODO: add mode examples
```

## Debug device

* see [Monitoring](https://github.com/opiproject/otel)
* uniform method of failure data collection and recovery
* explore SOL (Serial over LAN) as well

example if we using redfish:

```bash
$ curl -s -k -u <bmc-user>:<password> -X GET -H "Content-Type: application/json" https://<bmc-ip-address>/redfish/v1/Managers/<ID>/LogServices/{LogServiceId}/Entries
{}
$ curl -s -k -u <bmc-user>:<password> -X POST -H "Content-Type: application/json" https://<bmc-ip-address>/redfish/v1/Managers/<ID>/LogServices/{LogServiceId}/Actions/LogService.ClearLog
{}
# TODO: add more examples
```
