ARG GO_VER
ARG OS_VER

FROM golang:${GO_VER}${OS_VER} as builder

WORKDIR /opt

COPY ./cmd ./cmd
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go mod tidy && \
    go build -o /opt/build/transport /opt/cmd/stdlib/transport

FROM alpine:${OS_VER}

COPY --from=builder /opt/build/transport /usr/local/bin/transport

RUN apk update && \
    apk --no-cache add net-tools iproute2
