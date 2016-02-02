FROM ubuntu:14.04

ENV SHORT_VSN 2.0
ENV LONG_VSN 2.0.1
ENV TARGET_DEB riak-cs.deb

WORKDIR /
ADD http://s3.amazonaws.com/downloads.basho.com/riak-cs/${SHORT_VSN}/${LONG_VSN}/ubuntu/trusty/riak-cs_${LONG_VSN}-1_amd64.deb \
  ${TARGET_DEB}
RUN dpkg -i $TARGET_DEB
ENV NUM_PORTS 4096
RUN rm ${TARGET_DEB} && ulimit -n ${NUM_PORTS}

ADD bin /bin

CMD ["/bin/boot", "riak-cs"]
