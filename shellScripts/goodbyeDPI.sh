#!/bin/sh

/etc/dpitunnel-cli-arm64 --desync-attacks=fake,disorder_fake --split-position=2 --auto-ttl=1-4-10 --min-ttl=3 --doh --doh-server=https://dns.google/dns-query --wsize=1 --wsfactor=6 --ca-bundle-path=/etc/cacert.pem  --daemon --pid=/tmp/dpi.run --mode transparent --port 8081 &>/dev/null &

sysctl -w net.ipv4.ip_forward=1
sysctl -w net.ipv4.conf.all.send_redirects=0
iptables -t nat -A PREROUTING -i br-lan -p tcp --dport 80 -j REDIRECT --to-port 8081
iptables -t nat -A PREROUTING -i br-lan -p tcp --dport 443 -j REDIRECT --to-port 8081
