ARG GO_VER
ARG OS_VER

FROM golang:${GO_VER}${OS_VER} AS builder

WORKDIR /opt

COPY ./cmd ./cmd
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go build -o /opt/build/client /opt/cmd/p2p/client

FROM alpine:${OS_VER}

COPY --from=builder /opt/build/client /usr/local/bin/client

ENTRYPOINT [ "/usr/local/bin/client" ]
CMD "-port=2001"