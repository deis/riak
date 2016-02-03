FROM ubuntu:14.04

ENV RIAK_VERSION 2.0.2-1
ENV CONF_FOLDER /etc/riak
ENV CONF_FILE ${CONF_FOLDER}/riak.conf
ENV ADVANCED_CONF_FILE ${CONF_FOLDER}/advanced.config
ENV NUM_PORTS 65536

RUN \
    apt-get update && apt-get install -y curl dnsutils && \

    # Install Riak
    curl https://packagecloud.io/install/repositories/basho/riak/script.deb.sh | bash && \
    apt-get install -y riak=${RIAK_VERSION} && \

    # Cleanup
    apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Tune Riak configuration settings for the container
RUN sed -i.bak "s/storage_backend = \(.*\)//" ${CONF_FILE} && \
    echo "buckets.default.allow_mult = true" >> $CONF_FILE && \
    sed -i.bak 's/listener.http.internal = 127.0.0.1/listener.http.internal = 0.0.0.0/' ${CONF_FILE} && \
    sed -i.bak 's/listener.protobuf.internal = 127.0.0.1/listener.protobuf.internal = 0.0.0.0/' ${CONF_FILE} && \
    echo "anti_entropy.concurrency_limit = 1" >> ${CONF_FILE} && \
    echo "javascript.map_pool_size = 0" >> ${CONF_FILE} && \
    echo "javascript.reduce_pool_size = 0" >> ${CONF_FILE} && \
    echo "javascript.hook_pool_size = 0" >> ${CONF_FILE}
    # touch ${ADVANCED_CONF_FILE}
# this config file has a syntax error, omitting for now
# ADD etc/advanced.config ${ADVANCED_CONF_FILE}
ADD bin /bin

# Open ports for HTTP and Protocol Buffers
EXPOSE 8087 8098

CMD ["/bin/boot", "riak"]
