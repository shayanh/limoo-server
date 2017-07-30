FROM ubuntu:xenial

RUN apt-get update && apt-get install -y ca-certificates

EXPOSE 8000

ADD limood /limood
ADD config.yaml /config.yaml

ENTRYPOINT ["/limood"]