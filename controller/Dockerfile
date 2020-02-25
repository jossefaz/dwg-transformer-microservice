############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/package/app/
COPY ./ ./
COPY ../lib ./
COPY ../go.mod ./
# Fetch dependencies.
# Using go get.
RUN go get -d -v ./... && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/main
############################
# STEP 2 build a small image
############################
FROM alpine
# Copy our static executable.
WORKDIR /app

RUN mkdir shared config && \
    chmod 777 shared
COPY --from=builder /go/bin/main .
COPY ./config/*.json ./config/
# Run the hello binary.

CMD ["./main", "dev"]
