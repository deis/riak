#!/bin/bash

CONF_FILE=/etc/stanchion/stanchion.conf

if [ "$LISTEN_HOST" == "" ]; then
  echo "Error: LISTEN_HOST not set"
  exit 1
fi

if [ "$LISTEN_PORT" == "" ]; then
  echo "Error: LISTEN_PORT not set"
  exit 1
fi

ADMIN_KEY=$(cat /var/run/secrets/deis/riakcs/admin-user)
ADMIN_SECRET=$(cat /var/run/secrets/deis/riakcs/admin-secret)

sed -i.bak "s/listener = 127.0.0.1:8085/listener = $LISTEN_HOST:$LISTEN_PORT" $CONF_FILE
sed -i.bak "s/admin.key = admin-key/admin.key = $ADMIN_KEY" $CONF_FILE
sed -i.bak "s/admin.secret = admin-secret/admin.secret = $ADMIN_SECRET" $CONF_FILE
