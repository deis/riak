#!/bin/bash

CONF_FILE=/etc/riak-cs/riak-cs.conf
IP_ADDRESS=$(ip -o -4 addr list eth0 | awk '{print $4}' | cut -d/ -f1)
STANCHION_URL=${DEIS_RIAK_STANCHION_SERVICE_HOST}:${DEIS_RIAK_STANCHION_SERVICE_PORT}

if [ "$LISTEN_HOST" == "" ]; then
  echo "Error: LISTEN_HOST not set"
  exit 1
fi
if [ "$LISTEN_PORT" == "" ]; then
  echo "Error LISTEN_PORT not set"
  exit 1
fi
if [ "$STANCHION_URL" == "" ]; then
  echo "STANCHION_URL not set, exiting"
  exit 1
fi

ADMIN_KEY=$(cat /var/run/secrets/deis/riakcs/admin-user)
ADMIN_SECRET=$(cat /var/run/secrets/deis/riakcs/admin-secret)

sed -i.bak "s/listener = 127.0.0.1:8080/listener = $LISTEN_HOST:$LISTEN_PORT" $CONF_FILE
sed -i.bak \
  "s/stanchion_host = stanchion_host = 127.0.0.1:8085/stanchion_host = $STANCHION_URL" \
   $CONF_FILE
sed -i.bak "s/stanchion_ssl = on/stanchion_ssl = off" $CONF_FILE
sed -i.bak "s/admin.key = admin-key/admin.key = ${ADMIN_KEY}" $CONF_FILE
sed -i.bak "s/admin.secret = admin-secret/admin.secret = ${ADMIN_SECRET}" $CONF_FILE
