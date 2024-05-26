ARG GO_VER
ARG OS_VER

# Builder image
FROM golang:${GO_VER}${OS_VER} AS builder

WORKDIR /opt

COPY ./cmd ./cmd
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go build -o /opt/build/ping /opt/cmd/p2p/ping && \
    go build -o /opt/build/mdns /opt/cmd/p2p/mdns && \
    go build -o /opt/build/bootstrap /opt/cmd/p2p/bootstrap && \
    go build -o /opt/build/routing /opt/cmd/p2p/routing

# End images
FROM alpine:${OS_VER}

RUN apk update && apk add --no-cache bash

# Add the sysctl settings to a new file
RUN echo "net.core.rmem_max = 26214400" >> /etc/sysctl.conf && \
    echo "net.core.rmem_default = 26214400" >> /etc/sysctl.conf && \
    echo "net.core.wmem_max = 26214400" >> /etc/sysctl.conf && \
    echo "net.core.wmem_default = 26214400" >> /etc/sysctl.conf

COPY --from=builder /opt/build/ping /usr/local/bin/ping
COPY --from=builder /opt/build/mdns /usr/local/bin/mdns
COPY --from=builder /opt/build/bootstrap /usr/local/bin/bootstrap
COPY --from=builder /opt/build/routing /usr/local/bin/routing

