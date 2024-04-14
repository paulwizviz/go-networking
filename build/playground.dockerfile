ARG GO_VER
ARG OS_VER

FROM golang:${GO_VER}-bookworm AS builder

WORKDIR /opt

COPY ./cmd ./cmd
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go build -o /opt/build/selfaddr /opt/cmd/p2p/selfaddr

FROM ubuntu:${OS_VER}

COPY --from=builder /opt/build/selfaddr /usr/local/bin/selfaddr

RUN apt update && \
    apt install -y net-tools iproute2