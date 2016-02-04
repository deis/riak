FROM ubuntu:14.04

ENV TMP_DIR /tmp/stanchion
ENV LOG_DIR /var/log/stanchion
ENV NUM_PORTS 4096
ENV TARGET_DEB stanchion.deb
ENV CONF_FILE /etc/stanchion/stanchion.conf
RUN apt-get update && apt-get install -y curl dnsutils && \
    curl https://packagecloud.io/install/repositories/basho/stanchion/script.deb.sh | bash && \
    apt-get install -y stanchion=2.0.0-1 && \
    mkdir -p ${TMP_DIR} && \
    mkdir -p ${LOG_DIR}
ADD bin /bin
CMD ["/bin/boot", "riak-stanchion"]
