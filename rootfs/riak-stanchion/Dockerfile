FROM ubuntu:14.04

ENV TMP_DIR /tmp/riak-stanchion
ENV LOG_DIR /var/log/riak-stanchion
ENV NUM_PORTS 4096
ENV TARGET_DEB stanchion.deb
ADD http://s3.amazonaws.com/downloads.basho.com/stanchion/2.0/2.0.0/ubuntu/trusty/stanchion_2.0.0-1_amd64.deb ${TARGET_DEB}
RUN dpkg -i ${TARGET_DEB} && \
    rm ${TARGET_DEB} && \
    mkdir -p ${TMP_DIR} && \
    mkdir -p ${LOG_DIR}
ADD bin /bin
CMD ["/bin/boot", "riak-stanchion"]
