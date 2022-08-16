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

```text
$ docker-compose exec bootstrap curl --silent --fail -H Accept:application/yang-data+json http://localhost:1080/.well-known/host-meta
<XRD xmlns="http://docs.oasis-open.org/ns/xri/xrd-1.0">
  <Link rel="restconf" href="/restconf"/>
</XRD>
```

and also from another docker

```text
$ docker-compose exec agent curl --silent --fail -H Accept:application/yang-data+json http://bootstrap:1080/.well-known/host-meta
<XRD xmlns="http://docs.oasis-open.org/ns/xri/xrd-1.0">
  <Link rel="restconf" href="/restconf"/>
</XRD>
```

## More sZTP testing with simulator

See <https://watsen.net/support/sztpd-simulator-0.0.11.tgz>

```text
$ docker-compose exec agent bash

root@a204778c50cc:/tmp/sztpd-simulator# curl --fail -H Accept:application/yang-data+json http://bootstrap:1080/.well-known/host-meta
<XRD xmlns="http://docs.oasis-open.org/ns/xri/xrd-1.0">
  <Link rel="restconf" href="/restconf"/>
</XRD>

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
$ docker-compose run nmap
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
$ docker-compose run client
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
docker-compose run client curl --fail http://web:8082/var/lib/
```

OR

```text
docker run --network=sztp_opi --security-opt seccomp=unconfined -it --rm fedora:36 curl --fail http://web:8082/var/lib/
```
