ARG ARCH
FROM $ARCH/alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api

ARG BIN
COPY $BIN /

EXPOSE 80

ENTRYPOINT ["./app"]
