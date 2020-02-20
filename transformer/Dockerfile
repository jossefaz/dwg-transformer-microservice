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

FROM python:2.7.17-buster AS extracting
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
     ./configure --enable-trace --prefix=/usr/local && \
     make && \
     mkdir ldwg && \
     make install DESTDIR=$PWD/ldwg && \
     make check DESTDIR=$PWD/ldwg
     

FROM bitnami/minideb:jessie
COPY --from=extracting /app/libredwg/ldwg/usr/local/bin/* /usr/local/bin/
COPY --from=extracting /app/libredwg/ldwg/usr/local/include/* /usr/local/include/
COPY --from=extracting /app/libredwg/ldwg/usr/local/lib/* /usr/local/lib/
RUN ldconfig
COPY --from=builder /go/bin/main .
ENTRYPOINT [ "./main" ]