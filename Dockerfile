ARG ARCH
FROM $ARCH/alpine:latest
ARG BIN
COPY $BIN /app
EXPOSE 80
ENTRYPOINT ["./app"]
