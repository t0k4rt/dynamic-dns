#!/bin/bash
set -e
# for ubuntu
VERSION=1.0.0
OS=$(uname)
ARCH=$(uname -m)
USER=dynamicdns


while getopts ":v:" opt; do
  case $opt in
    v)
      echo "-v was triggered, Parameter: $OPTARG" >&2
      VERSION=$OPTARG
      ;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      exit 1
      ;;
    :)
      echo "Option -$OPTARG requires an argument." >&2
      exit 1
      ;;
  esac
done

echo "Install DynamiycDNS version $VERSION"

if id "$USER" &>/dev/null; then
  echo 'User already exists'
else
  if [[ "$OS" == "Linux" ]]; then
    useradd -r -s /bin/false "$USER"
  fi
fi

filename="dynamic-dns_${VERSION}_${OS}_${ARCH}.tar.gz"

SERVICE_FILE_URL="https://raw.githubusercontent.com/t0k4rt/dynamic-dns/${VERSION}/scripts/dynamicdns.service"
DEFAULT_CONFIG_FILE_URL="https://raw.githubusercontent.com/t0k4rt/dynamic-dns/${VERSION}/etc/dynamicdns.toml"
DEFAULT_ENV_FILE_URL="https://raw.githubusercontent.com/t0k4rt/dynamic-dns/${VERSION}/etc/dynamicdns_env.conf"

wget "https://github.com/t0k4rt/dynamic-dns/releases/download/${VERSION}/${filename}"
tar -xvf $filename
chmod +x dynamic-dns
mv dynamic-dns /usr/local/bin/

mkdir -p /etc/dynamicdns
curl -L "$DEFAULT_CONFIG_FILE_URL" -o /etc/dynamicdns/dynamicdns.toml
curl -L "$DEFAULT_ENV_FILE_URL" -o /etc/dynamicdns/dynamicdns_env.conf
chown -R "$USER:$USER" /etc/dynamicdns

touch /var/log/dynamicdns.log
chown "$USER:$USER" /var/log/dynamicdns.log

curl -L "$SERVICE_FILE_URL" -o /etc/systemd/system/dynamicdns.service

rm $filename
