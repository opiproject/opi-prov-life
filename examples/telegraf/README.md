# Telegraf

## Docs

* Plugins used
    * https://github.com/influxdata/telegraf/tree/master/plugins/inputs/http
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

Configuration file

```
$ cat telegraf-spdk.conf
[[inputs.http]]
  urls = ["http://127.0.0.1:9009"]
  headers = {"Content-Type" = "application/json"}
  method = "POST"
  username = "spdkuser"
  password = "spdkpass"
  body = '{"id":1, "method": "bdev_get_bdevs"}'
  data_format = "json"
  name_override = "dpu"
  json_strict = true
  tag_keys = ["name"]
  json_query = "result"

[[outputs.file]]
  files = ["stdout"]
  data_format = "influx"

[[outputs.influxdb_v2]]
  urls = ["http://127.0.0.1:8086"]
  token = "YourInfluxDBAuthToken"
  organization = "my-org"
  bucket = "my-bucket"
```

Run telegraf

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

dpu,host=localhost,name=Malloc0,url=http://127.0.0.1:9009 assigned_rate_limits_w_mbytes_per_sec=0,assigned_rate_limits_rw_ios_per_sec=0,assigned_rate_limits_r_mbytes_per_sec=0,block_size=512,num_blocks=131072,assigned_rate_limits_rw_mbytes_per_sec=0 1649268030000000000
dpu,host=localhost,name=Malloc1,url=http://127.0.0.1:9009 assigned_rate_limits_r_mbytes_per_sec=0,block_size=512,assigned_rate_limits_rw_ios_per_sec=0,assigned_rate_limits_rw_mbytes_per_sec=0,assigned_rate_limits_w_mbytes_per_sec=0,num_blocks=131072 1649268030000000000

...
```

## Grafana

see https://docs.influxdata.com/influxdb/v2.1/tools/grafana/

```
docker run --rm --net=host --name=grafana -p 3000:3000 grafana/grafana
```

Loging to http://127.0.0.1:3000/login using default (admin/admin)

Add new datasource InfluxDB with all the parameters above
