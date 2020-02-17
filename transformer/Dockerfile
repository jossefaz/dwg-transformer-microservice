############################
# STEP 1 build executable binary
############################


FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/package/app/
COPY . .
# Fetch dependencies.
# Using go get.
RUN go get -d -v
# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/main

FROM python:2.7.17-buster



RUN apt-get update && apt-get upgrade -y &&\
    apt-get install -y git autoconf libtool swig texinfo build-essential gcc python-libxml2 && \
    LIBXML2VER=2.9.1 && \
    mkdir libxmlInstall && cd libxmlInstall && \
    wget ftp://xmlsoft.org/libxml2/libxml2-$LIBXML2VER.tar.gz && \
    tar xf libxml2-$LIBXML2VER.tar.gz && \
    cd libxml2-$LIBXML2VER/ && \
    ./configure && \
    make && \
    make install && \
    cd /libxmlInstall && \
    rm -rf gg libxml2-$LIBXML2VER.tar.gz libxml2-$LIBXML2VER
WORKDIR /app
RUN git clone git://git.sv.gnu.org/libredwg.git && \
     cd libredwg && \
     sh autogen.sh && \
     ./configure --enable-trace && \
     make && \
     make install && \
     make check && \
     ldconfig

COPY --from=builder /go/bin/main .
EXPOSE 8090
ENTRYPOINT [ "./main" ] 