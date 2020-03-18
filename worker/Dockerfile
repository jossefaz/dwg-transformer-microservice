FROM golang:alpine AS gobuilder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/package/app/
COPY ./listener .
# Fetch dependencies.
# Using go get.
RUN go get -d -v ./... && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/main


FROM osgeo/gdal:alpine-normal-latest as geobuild

ENV PYTHON_PACKAGES="\
    pyproj Shapely pandas pika "

RUN apk add --no-cache --virtual build-dependencies \
    && apk add --virtual build-runtime build-base python3-dev openblas-dev freetype-dev pkgconfig gfortran \
    && ln -s /usr/include/locale.h /usr/include/xlocale.h \
    && python3 -m ensurepip \
    && rm -r /usr/lib/python*/ensurepip \
    && pip3 install --upgrade pip \
    && ln -sf /usr/bin/python3 /usr/bin/python \
    && ln -sf pip3 /usr/bin/pip \
    && rm -r /root/.cache \
    && pip install --no-cache-dir $PYTHON_PACKAGES \
    && apk add --no-cache --virtual build-dependencies $PACKAGES \
    && rm -rf /var/cache/apk/*

FROM osgeo/gdal:alpine-normal-latest

COPY --from=geobuild /usr/lib/python3.8 ./usr/lib/python3.8

WORKDIR /app

COPY ./utils ./utils
COPY ./resource ./resource
COPY ./registry ./registry
COPY ./Store ./Store
COPY --from=gobuilder /go/bin/main .
COPY ./listener/config/*.json ./config/

COPY bootstrap.py ./

RUN touch __init__.py

CMD ["./main", "prod"]

