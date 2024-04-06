#!/bin/bash

export WEBSERVER_IMAGE=go-networking/webserver:current
export HTTPUTIL_PROXY_IMAGE=go-networking/httputilproxy:current
export CUSTOM_PROXY_IMAGE=go-networking/customproxy:current
export P2P_NODE_IMAGE=go-networking/p2pnode:current

function build_image(){
    docker-compose -f ./build/builder.yml build
}

function clean_image(){
    docker rmi -f ${WEBSERVER_IMAGE}
    docker rmi -f ${HTTPUTIL_PROXY_IMAGE}
    docker rmi -f ${CUSTOM_PROXY_IMAGE}
    docker rmi -f $(docker images --filter "dangling=true" -q)
}