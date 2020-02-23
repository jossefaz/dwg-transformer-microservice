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
RUN cd src && go get -d -v && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/main
############################
# STEP 2 build a small image
############################
FROM alpine
# Copy our static executable.
WORKDIR /app
COPY --from=builder /go/bin/main .
RUN mkdir shared && chmod 777 shared

# Run the hello binary.

ENTRYPOINT ["./main"]