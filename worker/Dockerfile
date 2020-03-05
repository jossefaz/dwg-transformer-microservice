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


FROM osgeo/gdal:ubuntu-full-latest as venvbuilder
RUN apt update && \
    apt -y install python3-virtualenv

ENV VIRTUAL_ENV=./venv
RUN python3 -m virtualenv --python=/usr/bin/python3 $VIRTUAL_ENV
ENV PATH="$VIRTUAL_ENV/bin:$PATH"

# Install dependencies:
COPY requirements.txt .
RUN pip install -r requirements.txt

FROM osgeo/gdal:ubuntu-small-latest

WORKDIR /app

RUN apt update && \
    apt -y install python3-virtualenv

RUN mkdir -p shared/logs config && \
    chmod 777 shared

ENV VIRTUAL_ENV=./venv
RUN python3 -m virtualenv --python=/usr/bin/python3 $VIRTUAL_ENV
ENV PATH="$VIRTUAL_ENV/bin:$PATH"

COPY --from=venvbuilder ./venv ./venv

COPY --from=gobuilder /go/bin/main .
COPY ./listener/config/*.json ./config/

COPY bootstrap.py ./

CMD ["./main", "prod"]

