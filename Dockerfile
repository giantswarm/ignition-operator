FROM alpine:3.13.1

RUN apk add --no-cache ca-certificates

ADD ./ignition-operator /ignition-operator

ENTRYPOINT ["/ignition-operator"]
