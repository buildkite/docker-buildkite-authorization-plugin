FROM golang:1.9-alpine as builder
COPY . /go/src/github.com/buildkite/docker-buildkite-authorization-plugin
WORKDIR /go/src/github.com/buildkite/docker-buildkite-authorization-plugin
RUN set -ex \
    && apk add --no-cache --virtual .build-deps \
    gcc libc-dev \
    && go install -v -tags nosystemd -ldflags '-extldflags "-static"' \
    && apk del .build-deps
CMD ["/go/bin/docker-buildkite-authorization-plugin"]

FROM alpine
RUN apk update
RUN mkdir -p /run/docker/plugins /mnt/state /mnt/volumes
COPY --from=builder /go/bin/docker-buildkite-authorization-plugin .
CMD ["docker-buildkite-authorization-plugin"]
