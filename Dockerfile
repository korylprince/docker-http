FROM golang:1.10-alpine as builder

ARG CREDENTIALS
ARG VERSION

RUN apk add --no-cache git

RUN git clone --branch "$VERSION" --single-branch --depth 1 \
    https://github.com/korylprince/docker-http.git  /go/src/github.com/korylprince/docker-http

RUN go install github.com/korylprince/docker-http/serve

FROM alpine:3.7

COPY --from=builder /go/bin/serve /

RUN mkdir /http

CMD ["/serve", "-addr", ":80", "-dir", "/http"]
