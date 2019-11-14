ARG GO_VERSION=1.11

ARG ARCH
FROM $ARCH/golang:${GO_VERSION}-stretch AS builder

ARG QEMU_BIN
COPY $QEMU_BIN /usr/bin

RUN apt-get -y update && apt-get -y install git

RUN mkdir -p /api
WORKDIR /api

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./app ./main.go

ARG ARCH
FROM $ARCH/alpine:latest

ARG QEMU_BIN
COPY $QEMU_BIN /usr/bin

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api
COPY --from=builder /api/app .

EXPOSE 80

ENTRYPOINT ["./app"]
