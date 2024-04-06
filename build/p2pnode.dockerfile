ARG GO_VER
ARG OS_VER

FROM golang:${GO_VER}${OS_VER} AS builder

WORKDIR /opt

COPY ./cmd ./cmd
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go build -o /opt/build/p2p /opt/cmd/p2p

FROM alpine:${OS_VER}

COPY --from=builder /opt/build/p2p /usr/local/bin/p2p

CMD /usr/local/bin/p2p