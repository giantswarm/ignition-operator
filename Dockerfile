FROM alpine:3.13.4

RUN apk add --no-cache ca-certificates

ADD ./ignition-operator /ignition-operator

ENTRYPOINT ["/ignition-operator"]
