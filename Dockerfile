FROM alpine:3.7

COPY servidor /
ENTRYPOINT ["/servidor"]
