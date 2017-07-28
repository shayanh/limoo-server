FROM quay.io/brianredbeard/corebox

EXPOSE 8000

ADD limood /limood
ADD config.yaml /config.yaml

ENTRYPOINT ["/limood"]