ARG GO_VER
ARG OS_VER

FROM golang:${GO_VER}${OS_VER} AS builder

WORKDIR /opt

COPY ./cmd ./cmd
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go build -o /opt/build/ping /opt/cmd/p2p/ping

FROM alpine:${OS_VER}

COPY --from=builder /opt/build/ping /usr/local/bin/ping

ENTRYPOINT [ "/usr/local/bin/ping" ]
CMD "-port=2001"