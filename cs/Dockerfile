FROM ubuntu:14.04

RUN apt-get install -y curl apt-transport-https
RUN curl -s https://packagecloud.io/gpg.key | sudo apt-key add -

COPY ./install.sh /install-riakcs.sh
RUN /install-riakcs.sh
