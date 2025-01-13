## Project Name
pnsql-server

## Overview
pnsql 全能平台，由研发爱好者peng自研开发

## Web Installation
npm install

## Pnsql Installation
```
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
go mod tidy
go mod download
```

## config
```
mkdir /pnsql
cd /pnsql/go-project/PnSQL && go build -o pnsql main.go
sed -i "s|VITE_BASE_PATH= http://127.0.0.1:2379|VITE_BASE_PATH= http://8.153.100.186:2379|" .env.production
```

## Systemctl
```
sh -c 'cat <<EOF > /etc/systemd/system/pnsql-web.service
[Unit]
Description=PnSql Web Service
[Service]
ExecStart=/bin/bash -c "/root/.nvm/versions/node/v23.5.0/bin/npm run prod >> /var/log/pnsql-web.log 2>&1"
WorkingDirectory=/pnsql/go-project/PnSQL/web
Restart=always
RestartSec=5s
User=root
SyslogIdentifier=pnsql-web
[Install]
WantedBy=multi-user.target
EOF'
```

```
sh -c 'cat <<EOF > /etc/systemd/system/pnsql-server.service
[Unit]
Description=PnSql Server Service
[Service]
ExecStart=/bin/bash -c "./pnsql >> /var/log/pnsql-server.log 2>&1"
WorkingDirectory=/pnsql/go-project/PnSQL
Restart=always
RestartSec=5s
User=root
SyslogIdentifier=pnsql-server
[Install]
WantedBy=multi-user.target
EOF'
```

## run
```
systemctl daemon-reload
systemctl start pnsql-web.service
systemctl start pnsql-server.service
```
