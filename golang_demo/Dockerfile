FROM golang:1.6
MAINTAINER Mario Inga <mario21ic@gmail.com>

RUN wget https://raw.githubusercontent.com/pote/gpm/v1.4.0/bin/gpm && chmod +x gpm && mv gpm /usr/local/bin

ENV SRC_DIR=/app
RUN mkdir $SRC_DIR
COPY ./Godeps $SRC_DIR
WORKDIR $SRC_DIR
RUN gpm install

VOLUME $SRC_DIR
