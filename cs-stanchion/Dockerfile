FROM ubuntu:14.04

ENV SHORT_VSN 2.0
ENV LONG_VSN 2.0.0

ADD http://s3.amazonaws.com/downloads.basho.com/stanchion/${SHORT_VSN}/${LONG_VSN}/ubuntu/trusty/stanchion_${LONG_VSN}-1_amd64.deb /
RUN (cd / && dpkg -i "stanchion_${LONG_VSN}-1_amd64.deb")

RUN rm "/stanchion_${LONG_VSN}-1_amd64.deb"
