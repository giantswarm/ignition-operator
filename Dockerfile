FROM alpine:3.12.3

RUN apk add --no-cache ca-certificates

ADD ./ignition-operator /ignition-operator

ENTRYPOINT ["/ignition-operator"]
