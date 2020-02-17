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
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/main
############################
# STEP 2 build a small image
############################
FROM alpine
# Copy our static executable.
WORKDIR /app
COPY --from=builder /go/bin/main .
RUN mkdir uploads && chmod 777 uploads
# Run the hello binary.

EXPOSE 3030

ENTRYPOINT ["./main"]