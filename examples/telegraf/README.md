# Telegraf

## Docs

* Plugins used
    * https://github.com/influxdata/telegraf/tree/master/plugins/inputs/http
    * https://github.com/influxdata/telegraf/tree/master/plugins/inputs/cpu
    * https://github.com/influxdata/telegraf/tree/master/plugins/inputs/mem
    * https://github.com/influxdata/telegraf/tree/master/plugins/outputs/file
    * https://github.com/influxdata/telegraf/tree/master/plugins/outputs/influxdb_v2

* Dockers used
    * https://hub.docker.com/_/telegraf
    * https://hub.docker.com/_/influxdb
    * https://hub.docker.com/r/grafana/grafana
    * https://github.com/spdk/spdk-csi/blob/master/deploy/spdk/Dockerfile


## Getting started:

1. Run `docker-compose up -d`
2. Open `http://localhost:3000/explore` for querying InfluxDB

## SPDK RPC proxy

see https://spdk.io/doc/jsonrpc_proxy.html

Use this patch to handle chunked data https://review.spdk.io/gerrit/c/spdk/spdk/+/12082

```
sudo ./spdk/scripts/rpc_http_proxy.py 127.0.0.1 9009 spdkuser spdkpass
```

Test Proxy is running correctly
```
curl -k --user spdkuser:spdkpass -X POST -H "Content-Type: application/json" -d '{"id": 1, "method": "bdev_get_bdevs", "params": {"name": "Malloc0"}}' http://127.0.0.1:9009/
```

## InfluxDB

Use [docker-compose](docker-compose.yml) or manually

```
$ docker run --rm --name influxdb -p 8086:8086 \
      -e DOCKER_INFLUXDB_INIT_MODE=setup \
      -e DOCKER_INFLUXDB_INIT_USERNAME=my-user \
      -e DOCKER_INFLUXDB_INIT_PASSWORD=my-password \
      -e DOCKER_INFLUXDB_INIT_ORG=my-org \
      -e DOCKER_INFLUXDB_INIT_BUCKET=my-bucket \
      -e DOCKER_INFLUXDB_INIT_ADMIN_TOKEN=YourInfluxDBAuthToken \
      influxdb:2.1-alpine
```

## HTTP Input Plugin

Configuration file [telegraf-spdk.conf](config/telegraf-spdk.conf)

Run telegraf using [docker-compose](docker-compose.yml)

Example:
```
$ docker run --rm --net=host -v $(pwd)/telegraf-spdk.conf:/etc/telegraf/telegraf.conf:ro telegraf:1.22
2022-03-29T18:47:11Z I! Using config file: /etc/telegraf/telegraf.conf
2022-03-29T18:47:11Z I! Starting Telegraf 1.22.0
2022-03-29T18:47:11Z I! Loaded inputs: http
2022-03-29T18:47:11Z I! Loaded aggregators:
2022-03-29T18:47:11Z I! Loaded processors:
2022-03-29T18:47:11Z W! Outputs are not used in testing mode!
2022-03-29T18:47:11Z I! Tags enabled: host=localhost

dpu,host=localhost,name=Malloc0,url=http://127.0.0.1:9009 assigned_rate_limits_rw_mbytes_per_sec=0,num_blocks=131072,assigned_rate_limits_w_mbytes_per_sec=0,block_size=512,assigned_rate_limits_rw_ios_per_sec=0,assigned_rate_limits_r_mbytes_per_sec=0 1649268020000000000
dpu,host=localhost,name=Malloc1,url=http://127.0.0.1:9009 num_blocks=131072,assigned_rate_limits_w_mbytes_per_sec=0,assigned_rate_limits_rw_ios_per_sec=0,assigned_rate_limits_r_mbytes_per_sec=0,assigned_rate_limits_rw_mbytes_per_sec=0,block_size=512 1649268020000000000

mem,host=52ee5c75df01 commit_limit=69312983040i,committed_as=13494636544i,huge_page_size=2097152i,used_percent=10.100053796757296,high_free=0i,inactive=13541511168i,low_free=0i,shared=3904901120i,sreclaimable=812650496i,swap_cached=0i,free=100370612224i,huge_pages_total=2048i,low_total=0i,page_tables=49500160i,used=13567504384i,huge_pages_free=1357i,mapped=901996544i,slab=2243977216i,swap_total=4294963200i,cached=20385955840i,vmalloc_chunk=0i,vmalloc_used=0i,write_back=0i,swap_free=4294963200i,high_total=0i,available_percent=86.20598148102354,available=115801366528i,sunreclaim=1431326720i,total=134331011072i,buffered=6938624i,dirty=856064i,vmalloc_total=14073748835531776i,write_back_tmp=0i,active=8859537408i 1650954170000000000

net,host=52ee5c75df01,interface=eth0 drop_in=0i,drop_out=0i,bytes_sent=16589i,bytes_recv=13986i,packets_sent=89i,packets_recv=110i,err_in=0i,err_out=0i 1650954170000000000

cpu,cpu=cpu0,host=52ee5c75df01 usage_user=99.6999999973923,usage_system=0.09999999999763531,usage_idle=0,usage_iowait=0,usage_softirq=0,usage_steal=0,usage_nice=0,usage_irq=0.19999999999527063,usage_guest=0,usage_guest_nice=0 1650954170000000000
cpu,cpu=cpu1,host=52ee5c75df01 usage_user=99.70000000204891,usage_system=0,usage_irq=0.2999999999974534,usage_steal=0,usage_idle=0,usage_nice=0,usage_iowait=0,usage_softirq=0,usage_guest=0,usage_guest_nice=0 1650954170000000000
cpu,cpu=cpu2,host=52ee5c75df01 usage_system=0,usage_idle=0,usage_iowait=0,usage_guest_nice=0,usage_user=99.79999999981374,usage_nice=0,usage_irq=0.20000000000436557,usage_softirq=0,usage_steal=0,usage_guest=0 1650954170000000000
cpu,cpu=cpu3,host=52ee5c75df01 usage_guest_nice=0,usage_user=99.79999999981374,usage_idle=0,usage_nice=0,usage_iowait=0,usage_guest=0,usage_system=0,usage_irq=0.20000000000436557,usage_softirq=0,usage_steal=0 1650954170000000000
cpu,cpu=cpu4,host=52ee5c75df01 usage_user=99.70029970233988,usage_guest=0,usage_steal=0,usage_guest_nice=0,usage_system=0.09990009990223975,usage_idle=0,usage_nice=0,usage_iowait=0,usage_irq=0.19980019979993657,usage_softirq=0 1650954170000000000
cpu,cpu=cpu5,host=52ee5c75df01 usage_nice=0,usage_iowait=0,usage_irq=0.2997002997044478,usage_softirq=0,usage_steal=0,usage_guest_nice=0,usage_user=99.70029970233988,usage_idle=0,usage_guest=0,usage_system=0 1650954170000000000

...
```

## Grafana

see https://docs.influxdata.com/influxdb/v2.1/tools/grafana/

```
docker run --rm --net=host --name=grafana -p 3000:3000 grafana/grafana
```

Loging to http://127.0.0.1:3000/login using default (admin/admin)

Add new datasource InfluxDB with all the parameters above
