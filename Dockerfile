# First stage: build the Go application
FROM golang:1.20.3-alpine3.17 AS builder

RUN apk add -v build-base
RUN apk add -v ca-certificates
RUN apk add --no-cache \
    unzip \
    openssh

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

# RUN go build -o pb-build .

EXPOSE 8090

CMD ["go", "run", "main.go", "serve", "--http=127.0.0.1:8090"]
