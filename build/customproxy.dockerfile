FROM golang:1.22.0-alpine3.18 AS builder

WORKDIR /opt

COPY  ./cmd ./cmd
COPY ./internal ./internal
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go build -o /opt/build/custom /opt/cmd/proxy/custom

FROM alpine:3.18

COPY --from=builder /opt/build/custom /usr/local/bin/custom
CMD /usr/local/bin/custom