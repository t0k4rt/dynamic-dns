#!/bin/bash

# for ubuntu
latest_tag=$(git describe --abbrev=0 --tags)
os=$(uname)
arch=$(uname -m)
filename="dynamic-dns_${latest_tag}_${os}_${arch}.tar.gz"

wget "https://github.com/t0k4rt/dynamic-dns/releases/download/${latest_tag}/${filename}"
tar -xvf $filename
chmod +x dynamic-dns
mv dynamic-dns /usr/local/bin/

mkdir -p /etc/dynamicdns
cp ./etc/dynamicdns.toml /etc/dynamicdns/dynamicdns.toml
cp ./etc/dynamicdns_env.conf /etc/dynamicdns/dynamicdns_env.conf
chown -R www-data:www-data /etc/dynamicdns

touch /var/log/dynamicdns.log
chown www-data:www-data /var/log/dynamicdns.log

cp ./scripts/dynamicdns.service /etc/systemd/system/dynamicdns.service
rm $filename