ARG GO_VER
ARG OS_VER

# Builder image
FROM golang:${GO_VER}${OS_VER} AS builder

WORKDIR /opt

COPY ./cmd ./cmd
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go build -o /opt/build/ping /opt/cmd/p2p/ping && \
    go build -o /opt/build/mdns /opt/cmd/p2p/mdns

# End images
FROM alpine:${OS_VER}

COPY --from=builder /opt/build/ping /usr/local/bin/ping
COPY --from=builder /opt/build/mdns /usr/local/bin/mdns

