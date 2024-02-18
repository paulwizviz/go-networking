FROM golang:1.22.0-alpine3.18 AS builder

WORKDIR /opt

COPY  ./cmd ./cmd
COPY ./internal ./internal
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go build -o /opt/build/customproxy /opt/cmd/customproxy

FROM alpine:3.18

COPY --from=builder /opt/build/customproxy /usr/local/bin/customproxy
CMD /usr/local/bin/customproxy