FROM golang:1.22.0-alpine3.18 AS builder

WORKDIR /opt

COPY  ./cmd ./cmd
COPY ./internal ./internal
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go build -o /opt/build/httputil /opt/cmd/proxy/httputil

FROM alpine:3.18

COPY --from=builder /opt/build/httputil /usr/local/bin/httputil
CMD /usr/local/bin/httputil