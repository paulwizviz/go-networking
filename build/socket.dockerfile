ARG GO_VER
ARG OS_VER

FROM golang:${GO_VER}${OS_VER} AS builder

WORKDIR /opt

COPY  ./cmd ./cmd
COPY ./internal ./internal
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go build -o /usr/local/bin/socket /opt/cmd/socket