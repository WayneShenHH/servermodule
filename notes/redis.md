# redis

install redis-cli
>sudo apt-get install redis-tools

config redis cluster

```shell
# cluster meet
redis-cli -p 7001 cluster meet 127.0.0.1 7002
redis-cli -p 7001 cluster meet 127.0.0.1 7003
redis-cli -p 7001 cluster meet 127.0.0.1 7004
redis-cli -p 7001 cluster meet 127.0.0.1 7005
redis-cli -p 7001 cluster meet 127.0.0.1 7006
# cluster addslots
redis-cli -p 7001 cluster addslots {0..5641}
redis-cli -p 7002 cluster addslots {5642..10922}
redis-cli -p 7003 cluster addslots {10923..16383}
# looking hash id
redis-cli -p 7001 cluster nodes
# cluster replicate (7004 replicate 7001, 7005 replicate 7002, 7006 replicate 7003)
redis-cli -p 7004 cluster replicate 7386de871d5f7be4848bfeeacd8f5667b4a43bef
redis-cli -p 7005 cluster replicate 4231cb4151a24d8ee40b0156e6487136f8f00ccb
redis-cli -p 7006 cluster replicate 592e95d6aab1326ae82cc73e1402b3a33acad641
# 3 master & 3 slave
cluster nodes
```
