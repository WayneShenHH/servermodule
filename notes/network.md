connect ws

```
curl -i -N \
-H "Connection: Upgrade" \
-H "Upgrade: websocket" \
-H "Sec-Websocket-Version: 13" \
-H 'Sec-WebSocket-Key: +onQ3ZxjWlkNa0na6ydhNg==' \
http://localhost:8899/echo
```

tcpdump ('lo' is local network, find it by $ ip a)

```
sudo tcpdump -v -w /tmp/tcp_http.pcap -i lo port 18086
```

Netcat to ping a host

```
nc -zv localhost 18086
```

Wireshark read tcpdump file

```
wireshark /tmp/tcp_http.pcap
```

show ip infos
```
sudo lshw -class network -short
ip a
```

ping reverse-proxy.sit-gm.svc.cluster.local issue
```
systemctl stop avahi-daemon.service
systemctl disable avahi-daemon.service

systemctl stop avahi-daemon.socket
systemctl disable avahi-daemon.socket
```