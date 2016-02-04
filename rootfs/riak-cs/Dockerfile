FROM ubuntu:14.04

ENV TARGET_DEB riak-cs.deb
ENV CONF_FOLDER /etc/riak-cs
ENV CONF_FILE ${CONF_FOLDER}/riak-cs.conf
ENV NUM_PORTS 4096
ENV TMP_DIR /tmp/riak-cs
ENV LOG_DIR /var/log/riak-cs

RUN apt-get update && apt-get install -y curl dnsutils && \
    curl https://packagecloud.io/install/repositories/basho/riak-cs/script.deb.sh | bash && \
    apt-get install -y riak-cs=2.0.0-1 && \
    apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* && \
    sed -i.bak "s/stanchion_ssl = on/stanchion_ssl = off/" ${CONF_FILE} && \
    mkdir -p ${TMP_DIR} && \
    mkdir -p ${LOG_DIR}
ADD bin /bin
CMD ["/bin/boot", "riak-cs"]
