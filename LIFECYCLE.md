# Lifecycle Management

High level requirements on a DPU/IPU or subsequently a "device"

## Reboot device

example if we using redfish:
```bash
$ curl -s -k -u <bmc-user>:<password> -X POST -H "Content-Type: application/json" -d '{"ResetType": "ForceRestart"}' https://<bmc-ip-address>/redfish/v1/Managers/<ID>/Actions/Manager.Reset
and
$ curl -s -k -u <bmc-user>:<password> -X POST -H "Content-Type: application/json" -d '{"ResetType": "PowerCycle"}' https://<bmc-ip-address>/redfish/v1/Systems/<ID>/Actions/ComputerSystem.Reset 
```

## Update device

 * firmware update
 * OS update
 * software/application update

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
```

## Debug device

* see [Monitoring](MONITORING.md)
* uniform method of failure data collection and recovery
* explore SOL (Serial over LAN) as well
