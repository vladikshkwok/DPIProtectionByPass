#!/bin/sh

#me of the ipset
IPSET_NAME="youtube_ipset"

# Proxy port
PROXY_PORT=8081

# File with domain names
DOMAINS_FILE="domains.txt"

# Path to the dnsmasq configuration file
DNSMASQ_CONF="/etc/dnsmasq.d/ipset.conf"

# Check if the domains file exists
if [ ! -f "$DOMAINS_FILE" ]; then
  echo "File $DOMAINS_FILE not found!"
  exit 1
fi

# Create the ipset if it does not exist
if ! ipset list $IPSET_NAME >/dev/null 2>&1; then
    ipset create $IPSET_NAME hash:ip
fi

# Flush the ipset
ipset flush $IPSET_NAME

# Create the dnsmasq configuration
echo "server=8.8.8.8" > $DNSMASQ_CONF
echo "server=8.8.4.4" >> $DNSMASQ_CONF

while IFS= read -r domain; do
  echo "ipset=/$domain/$IPSET_NAME" >> $DNSMASQ_CONF
done < "$DOMAINS_FILE"

# Restart dnsmasq
/etc/init.d/dnsmasq restart


# Remove old iptables rules if they exist
iptables -t nat -D PREROUTING -p tcp -m set --match-set $IPSET_NAME dst -j REDIRECT --to-ports $PROXY_PORT 2>/dev/null
iptables -t nat -D PREROUTING -p udp -m set --match-set $IPSET_NAME dst -j REDIRECT --to-ports $PROXY_PORT 2>/dev/null

# Add new iptables rules to redirect traffic through the proxy
iptables -t nat -A PREROUTING -p tcp -m set --match-set $IPSET_NAME dst -j REDIRECT --to-ports $PROXY_PORT
iptables -t nat -A PREROUTING -p udp -m set --match-set $IPSET_NAME dst -j REDIRECT --to-ports $PROXY_PORT

echo "Script executed successfully"