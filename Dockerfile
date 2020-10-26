FROM alpine:3.12.1

RUN apk add --no-cache ca-certificates

ADD ./ignition-operator /ignition-operator

ENTRYPOINT ["/ignition-operator"]
