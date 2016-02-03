FROM ubuntu:14.04

ENV TARGET_DEB riak-cs.deb
ENV CONF_FOLDER /etc/riak-cs
ENV CONF_FILE ${CONF_FOLDER}/riak-cs.conf
ENV NUM_PORTS 4096
ADD http://s3.amazonaws.com/downloads.basho.com/riak-cs/2.0/2.0.0/ubuntu/trusty/riak-cs_2.0.0-1_amd64.deb ${TARGET_DEB}
RUN dpkg -i $TARGET_DEB
RUN rm ${TARGET_DEB} && \
    ulimit -n ${NUM_PORTS} && \
    sed -i.bak "s/stanchion_ssl = on/stanchion_ssl = off/" ${CONF_FILE}
ADD bin /bin
CMD ["/bin/boot", "riak-cs"]
