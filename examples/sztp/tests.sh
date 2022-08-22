#!/bin/bash
# SPDX-License-Identifier: Apache-2.0
# Copyright (c) 2022 Dell Inc, or its subsidiaries.

set -euxo pipefail

# let everything start
sleep 5

# test dhcp server
docker-compose exec -T dhcp cat /var/lib/dhcpd/dhcpd.leases

# let server respond
sleep 5

# tests dhcp client
docker-compose exec -T client cat /var/lib/dhclient/dhclient.leases
docker-compose exec -T client cat /var/lib/dhclient/dhclient.leases | grep sztp-redirect-urls
BOOTSTRAP=$(docker-compose exec -T client cat /var/lib/dhclient/dhclient.leases | grep sztp-redirect-urls | head -n 1 | awk '{print $3}' | tr -d '";')

# send server configuration
docker-compose exec -T bootstrap curl -i -X PUT --user my-admin@example.com:my-secret --data @/tmp/running.json -H "Content-Type:application/yang-data+json" http:/bootstrap:1080/restconf/ds/ietf-datastores:running

# let server reboot
sleep 5

# read back to check configuration was set
docker-compose exec -T bootstrap curl -i --user my-admin@example.com:my-secret -H "Accept:application/yang-data+json" http://bootstrap:1080/restconf/ds/ietf-datastores:running

# request onboarding info (like a DPU or IPU device would)
docker-compose run --rm -T agent curl --silent X POST --data @/tmp/input.json -H "Content-Type:application/yang-data+json" --user my-serial-number:my-secret http://bootstrap:9090/restconf/operations/ietf-sztp-bootstrap-server:get-bootstrapping-data | tee /tmp/post_rpc_input.json
docker-compose run --rm -T agent curl --silent X POST --data @/tmp/input.json -H "Content-Type:application/yang-data+json" --user my-serial-number:my-secret "${BOOTSTRAP}" | tee /tmp/post_rpc_input.json

# parse the reply
jq -r .\"ietf-sztp-bootstrap-server:output\".\"conveyed-information\" /tmp/post_rpc_input.json | base64 --decode | tail -n +2 | sed  '1i {' | jq . | tee /tmp/post_rpc_fixed.json

# check audit log
docker-compose exec -T bootstrap curl -i -X GET --user my-admin@example.com:my-secret  -H "Accept:application/yang-data+json" http://bootstrap:1080/restconf/ds/ietf-datastores:operational/wn-sztpd-1:audit-log

# parse the reply some more
jq -r .\"ietf-sztp-conveyed-info:onboarding-information\".\"configuration\" /tmp/post_rpc_fixed.json | base64 --decode

jq -r .\"ietf-sztp-conveyed-info:onboarding-information\".\"pre-configuration-script\" /tmp/post_rpc_fixed.json | base64 --decode

jq -r .\"ietf-sztp-conveyed-info:onboarding-information\".\"post-configuration-script\" /tmp/post_rpc_fixed.json | base64 --decode

jq -r .\"ietf-sztp-conveyed-info:onboarding-information\".\"boot-image\".\"download-uri\"[] /tmp/post_rpc_fixed.json

jq -r .\"ietf-sztp-conveyed-info:onboarding-information\".\"boot-image\".\"image-verification\"[] /tmp/post_rpc_fixed.json

docker-compose run --rm -T agent curl --fail http://web:8082/var/lib/misc/

# actually go and download the image from the web server
URL=$(jq -r .\"ietf-sztp-conveyed-info:onboarding-information\".\"boot-image\".\"download-uri\"[] /tmp/post_rpc_fixed.json)
docker-compose run --rm -v /tmp:/tmp agent curl --output /tmp/"$(basename "${URL}")" --fail "${URL}"

# Validate signature
SIGNATURE=$(openssl dgst -sha256 -c /tmp/"$(basename "${URL}")" | awk '{print $2}')
jq -r .\"ietf-sztp-conveyed-info:onboarding-information\".\"boot-image\".\"image-verification\"[] /tmp/post_rpc_fixed.json | grep "${SIGNATURE}"

echo "DONE"
