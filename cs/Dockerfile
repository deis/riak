FROM ubuntu:14.04

ENV SHORT_VSN 1.5
ENV LONG_VSN 1.5.2

ADD http://s3.amazonaws.com/downloads.basho.com/riak-cs/${SHORT_VSN}/${LONG_VSN}/ubuntu/trusty/riak-cs_${LONG_VSN}-1_amd64.deb /
RUN (cd / && dpkg -i "riak-cs_${LONG_VSN}-1_amd64.deb")

RUN rm "/riak-cs_${LONG_VSN}-1_amd64.deb"
