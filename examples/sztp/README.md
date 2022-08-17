# sZTP and DHCP adn HTTP

## Docs

* <https://www.rfc-editor.org/rfc/pdfrfc/rfc8572.txt.pdf> (Secure Zero Touch Provisioning (SZTP))
* <https://watsen.net/docs/sztpd/0.0.11/admin-guide/#simulator>

## sZTP Diagram

![xPU sZTP provisioning participants](../../architecture/sZTP-components.png)

## sZTP on DPU Diagram

![xPU sZTP provisioning block](../../architecture/sZTP-provisioning-blocks.png)

## Run sZTP

```text
docker-compose up --build bootstrap
```

## Test sZTP

Fetching Host-meta

```text
$ docker-compose run --rm -T agent curl -i --fail -H Accept:application/yang-data+json http://bootstrap:1080/.well-known/host-meta
HTTP/1.1 200 OK
Content-Type: application/xrd+xml; charset=utf-8
Content-Length: 104
Date: Wed, 17 Aug 2022 00:29:54 GMT
Server: <redacted>

<XRD xmlns="http://docs.oasis-open.org/ns/xri/xrd-1.0">
  <Link rel="restconf" href="/restconf"/>
</XRD>
```

Fetching the RESTCONF Root Resource

```text
$ docker-compose run --rm -T agent curl -i --fail -H Accept:application/yang-data+json http://bootstrap:1080/restconf/
HTTP/1.1 200 OK
Content-Type: application/yang-data+json; charset=utf-8
Content-Length: 137
Date: Wed, 17 Aug 2022 00:30:32 GMT
Server: <redacted>

{
    "ietf-restconf:restconf" : {
        "data" : {},
        "operations" : {},
        "yang-library-version" : "2019-01-04"
    }
}
```

Get the Current (Default) Configuration

```text
$ docker-compose run --rm -T agent curl -i -H "Accept:application/yang-data+json" http://bootstrap:1080/restconf/ds/ietf-datastores:running
HTTP/1.1 200 OK
Content-Type: application/yang-data+json; charset=utf-8
Content-Length: 318
Date: Wed, 17 Aug 2022 00:24:47 GMT
Server: <redacted>

{
  "wn-sztpd-1:transport": {
    "listen": {
      "endpoint": [
        {
          "name": "default startup endpoint",
          "use-for": "native-interface",
          "http": {
            "tcp-server-parameters": {
              "local-address": "0.0.0.0"
            }
          }
        }
      ]
    }
  }
}
```

## Device Getting Onboarding Information

Send the configuration file to SZTPD

```text
$ docker-compose exec bootstrap curl -i -X PUT --user my-admin@example.com:my-secret --data @/tmp/running.json -H "Content-Type:application/yang-data+json" http:/bootstrap:1080/restconf/ds/ietf-datastores:running

HTTP/1.1 204 No Content
Date: Wed, 17 Aug 2022 19:15:57 GMT
Server: <redacted>
```

Read the configuration back and validate it is correct:

```text
docker-compose exec bootstrap curl -i --user my-admin@example.com:my-secret -H "Accept:application/yang-data+json" http://bootstrap:1080/restconf/ds/ietf-datastores:running
```

Get onboarding info (from device perspective)

```text
$ docker-compose run --rm -T agent curl --silent X POST --data @/tmp/input.json -H "Content-Type:application/yang-data+json" --user my-serial-number:my-secret http://bootstrap:9090/restconf/operations/ietf-sztp-bootstrap-server:get-bootstrapping-data | tee /tmp/post_rpc_input.json

Creating sztp_agent_run ... done
{
  "ietf-sztp-bootstrap-server:output": {
    "conveyed-information": "MIIDOQYLKoZIhvcNAQkQASugggMoBIIDJHsKICAiaWV0Zi1zenRwLWNvbnZleWVkLWluZm86b25ib2FyZGluZy1pbmZvcm1hdGlvbiI6IHsKICAgICJib290LWltYWdlIjogewogICAgICAiZG93bmxvYWQtdXJpIjogWwogICAgICAgICJodHRwczovL2V4YW1wbGUuY29tL215LWJvb3QtaW1hZ2UuaW1nIgogICAgICBdLAogICAgICAiaW1hZ2UtdmVyaWZpY2F0aW9uIjogWwogICAgICAgIHsKICAgICAgICAgICJoYXNoLWFsZ29yaXRobSI6ICJpZXRmLXN6dHAtY29udmV5ZWQtaW5mbzpzaGEtMjU2IiwKICAgICAgICAgICJoYXNoLXZhbHVlIjogIjdiOmNhOmU2OmFjOjIzOjA2OmQ4Ojc5OjA2OjhjOmFjOjAzOjgwOmUyOjE2OjQ0OjdlOjQwOjZhOjY1OmZhOmQ0OjY5OjYxOjZlOjA1OmNlOmY1Ojg3OmRjOjJiOjk3IgogICAgICAgIH0KICAgICAgXQogICAgfSwKICAgICJwcmUtY29uZmlndXJhdGlvbi1zY3JpcHQiOiAiSXk5aWFXNHZZbUZ6YUFwbFkyaHZJQ0pwYm5OcFpHVWdkR2hsSUhCeVpTMWpiMjVtYVdkMWNtRjBhVzl1TFhOamNtbHdkQzR1TGlJSyIsCiAgICAiY29uZmlndXJhdGlvbi1oYW5kbGluZyI6ICJtZXJnZSIsCiAgICAiY29uZmlndXJhdGlvbiI6ICJQSFJ2Y0NCNGJXeHVjejBpYUhSMGNITTZMMlY0WVcxd2JHVXVZMjl0TDJOdmJtWnBaeUkrQ2lBZ1BHRnVlUzE0Yld3dFkyOXVkR1Z1ZEMxdmEyRjVMejRLUEM5MGIzQStDZz09IiwKICAgICJwb3N0LWNvbmZpZ3VyYXRpb24tc2NyaXB0IjogIkl5OWlhVzR2WW1GemFBcGxZMmh2SUNKcGJuTnBaR1VnZEdobElIQnZjM1F0WTI5dVptbG5kWEpoZEdsdmJpMXpZM0pwY0hRdUxpNGlDZz09IgogIH0KfQ=="
  }
}
```

Decode payload

```text
$ jq -r .\"ietf-sztp-bootstrap-server:output\".\"conveyed-information\" /tmp/post_rpc_input.json | base64 --decode
0▒9
   *▒H▒▒+▒▒(▒${
  "ietf-sztp-conveyed-info:onboarding-information": {
    "boot-image": {
      "download-uri": [
        "https://example.com/my-boot-image.img"
      ],
      "image-verification": [
        {
          "hash-algorithm": "ietf-sztp-conveyed-info:sha-256",
          "hash-value": "7b:ca:e6:ac:23:06:d8:79:06:8c:ac:03:80:e2:16:44:7e:40:6a:65:fa:d4:69:61:6e:05:ce:f5:87:dc:2b:97"
        }
      ]
    },
    "pre-configuration-script": "Iy9iaW4vYmFzaAplY2hvICJpbnNpZGUgdGhlIHByZS1jb25maWd1cmF0aW9uLXNjcmlwdC4uLiIK",
    "configuration-handling": "merge",
    "configuration": "PHRvcCB4bWxucz0iaHR0cHM6L2V4YW1wbGUuY29tL2NvbmZpZyI+CiAgPGFueS14bWwtY29udGVudC1va2F5Lz4KPC90b3A+Cg==",
    "post-configuration-script": "Iy9iaW4vYmFzaAplY2hvICJpbnNpZGUgdGhlIHBvc3QtY29uZmlndXJhdGlvbi1zY3JpcHQuLi4iCg=="
  }
}
```

View the Audit Log

```text
$ docker-compose exec bootstrap curl -i -X GET --user my-admin@example.com:my-secret  -H "Accept:application/yang-data+json" http://bootstrap:1080/restconf/ds/ietf-datastores:operational/wn-sztpd-1:audit-log

HTTP/1.1 200 OK
Content-Type: application/yang-data+json; charset=utf-8
Content-Length: 648
Date: Wed, 17 Aug 2022 19:35:34 GMT
Server: <redacted>

{
  "wn-sztpd-1:audit-log": {
    "log-entry": [
      {
        "timestamp": "2022-08-17T19:35:22Z",
        "source-ip": "10.127.127.3",
        "host": "bootstrap:9090",
        "method": "POST",
        "path": "/restconf/operations/ietf-sztp-bootstrap-server:get-bootstrapping-data",
        "outcome": "success"
      }
    ]
  }
}
```

## More sZTP testing with simulator

See <https://watsen.net/support/sztpd-simulator-0.0.11.tgz>

```text
$ docker-compose run --rm -T agent bash
root@a204778c50cc:/tmp/sztpd-simulator#
```

Provision certificates and run agent

```text
root@a204778c50cc:/tmp/sztpd-simulator# cat pki/sztpd1/sbi/root-ca/my_cert.pem pki/sztpd1/sbi/intermediate1/my_cert.pem  > /tmp/trust_chain.pem

root@a204778c50cc:/tmp/sztpd-simulator# ./rfc8572-agent.sh my-serial-number my-secret pki/client/end-entity/private_key.pem pki/client/end-entity/my_cert.pem bootstrap 1080 /tmp/trust_chain.pem
WARNING: Package(s) not found: sztpd
  ^-- Getting bootstrapping data...
failed (incorrect curl exit status code) on line 437.
  - exit code: 35 (expected 0)
  - command: curl --silent --dump-header header.dump -H Accept:application/yang-data+json --cacert /tmp/trust_chain.pem --key pki/client/end-entity/private_key.pem --cert pki/client/end-entity/my_cert.pem --user my-serial-number:my-secret https://bootstrap:1080/.well-known/host-meta
  - output: <starts on next line>
```

## Run DHCP

```text
docker-compose up --build dhcp
```

## Test DHCP with NMAP

```text
$ docker-compose run --rm -T nmap
Creating sztp_nmap_run ... done
Starting Nmap 7.92 ( https://nmap.org ) at 2022-08-15 19:13 UTC
Pre-scan script results:
| broadcast-dhcp-discover:
|   Response 1 of 1:
|     Interface: eth0
|     IP Offered: 10.127.127.101
|     DHCP Message Type: DHCPOFFER
|     Server Identifier: 10.127.127.2
|     IP Address Lease Time: 10m00s
|     Subnet Mask: 255.255.255.0
|     Bootfile Name: test.cfg
|_    TFTP Server Name: w.x.y.z
WARNING: No targets were specified, so 0 hosts scanned.
Nmap done: 0 IP addresses (0 hosts up) scanned in 10.25 seconds
```

## Test DHCP with client

```text
$ docker-compose run --rm -T client
Creating sztp_client_run ... done
Internet Systems Consortium DHCP Client 4.4.3
Copyright 2004-2022 Internet Systems Consortium.
All rights reserved.
For info, please visit https://www.isc.org/software/dhcp/

RTNETLINK answers: Operation not permitted
Listening on LPF/eth0/02:42:0a:7f:7f:03
Sending on   LPF/eth0/02:42:0a:7f:7f:03
Sending on   Socket/fallback
DHCPDISCOVER on eth0 to 255.255.255.255 port 67 interval 8 (xid=0xb3f32238)
DHCPOFFER of 10.127.127.102 from 10.127.127.2
DHCPREQUEST for 10.127.127.102 on eth0 to 255.255.255.255 port 67 (xid=0xb3f32238)
DHCPACK of 10.127.127.102 from 10.127.127.2 (xid=0xb3f32238)
RTNETLINK answers: Operation not permitted
bound to 10.127.127.102 -- renewal in 263 seconds.
```

see result

```text
$ docker-compose exec client cat /var/lib/dhclient/dhclient.leases
lease {
  interface "eth0";
  fixed-address 10.127.127.100;
  filename "grubx64.efi";
  option subnet-mask 255.255.255.0;
  option sztp-redirect-urls "http://192.0.2.1/demo.sh";
  option dhcp-lease-time 600;
  option tftp-server-name "w.x.y.z";
  option bootfile-name "test.cfg";
  option dhcp-message-type 5;
  option dhcp-server-identifier 10.127.127.2;
  renew 1 2022/08/15 19:16:40;
  rebind 1 2022/08/15 19:20:50;
  expire 1 2022/08/15 19:22:05;
}
```

## Run HTTP

```text
docker-compose up --build web
```

## Test HTTP from client

```text
docker-compose run --rm -T client curl --fail http://web:8082/var/lib/
```

OR

```text
docker run --network=sztp_opi --security-opt seccomp=unconfined -it --rm fedora:36 curl --fail http://web:8082/var/lib/
```
