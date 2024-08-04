Несколько скриптов позволяющих используя [dpitunnel-cli](https://github.com/nomoresat/DPITunnel-cli) обходить DPI. 
Также небольшой веб-сервер на go со свитчером этой защиты.

На роутере с openWRT (либо с доступом по ssh) можно прописать в крон запуск googbyeDPI.sh. 
Веб сервер же надо сначала сбилдить (env GOOS=linux GOARCH=arm64 go build -o forarm64) и закинуть его в /etc (а также закинуть туда views, js, css директории).
В файле domains.txt лежат доменные имена для обхода DPI.


На роутере должны быть доступны: dnsmasq, ipset, iptables


Т.к. на скорую руку, везде по путям и портам хардкод :(
