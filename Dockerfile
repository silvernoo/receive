ARG GO_VERSION=1.11

ARG ARCH
FROM $ARCH/golang:${GO_VERSION}-alpine AS builder

ARG QEMU_BIN
COPY $QEMU_BIN /usr/bin

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api

COPY go.mod .
COPY go.sum .

RUN export GOPROXY="https://goproxy.io" \
    && export CGO_ENABLED="0" \
    && go mod download

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
