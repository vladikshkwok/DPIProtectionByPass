#!/bin/sh

kill $(pgrep DPITunnel)
delete=0

while [[ "$delete" == "0" ]]; do
  echo "Delete iptables rules"
  iptables -t nat -D PREROUTING 4
  iptables -t nat -D PREROUTING 4
  delete=$?
done
