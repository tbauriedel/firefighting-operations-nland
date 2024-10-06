# fire-operations-nland

Scrapes firefighting operations from kfv-online.de and sends them to telegram via api

## Installation

```shell
groupadd firefighting-operations-nland
useradd -r -g firefighting-operations-nland -d /etc/firefighting-operations-nland -s /sbin/nologin firefighting-operations-nland
install -d -o firefighting-operations-nland -g firefighting-operations-nland -m 2750 /etc/firefighting-operations-nland

git clone https://github.com/tbauriedel/firefighting-operations-nland
cd firefighting-operations-nland
go build .
cp firefighting-operations-nland /usr/local/bin/firefighting-operations-nland
cp contrib/systemd/firefighting-operations-nland.service /etc/systemd/system
systemctl daemon-reload
```

## Quick start
Generate example config with `firefighting-operations-nland --generate-config` and add your telegram bot / channel details.
Start service `systemctl start firefighting-operations-nland.service`

