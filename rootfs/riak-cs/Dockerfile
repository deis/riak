FROM ubuntu:14.04

ENV TARGET_DEB riak-cs.deb
ENV NUM_PORTS 4096
ADD http://s3.amazonaws.com/downloads.basho.com/riak-cs/2.0/2.0.0/ubuntu/trusty/riak-cs_2.0.0-1_amd64.deb ${TARGET_DEB}
RUN dpkg -i $TARGET_DEB
RUN rm ${TARGET_DEB} && ulimit -n ${NUM_PORTS}
ADD bin /bin
CMD ["/bin/boot", "riak-cs"]
