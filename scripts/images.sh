#!/bin/bash

# Playground
export PLAYGROUND_IMAGE=go-networking/playground:current
function build_playground(){
    docker-compose -f ./build/builder.yml build playground
}

function clean_playground(){
    docker rmi -f ${PLAYGROUND_IMAGE}
    docker rmi -f $(docker images --filter "dangling=true" -q)
}

# P2P
export P2P_NODE_IMAGE=go-networking/p2pnode:current
function build_p2pnode(){
    docker-compose -f ./build/builder.yml build p2pnode
}

function clean_p2pnode(){
    docker rmi -f ${P2P_NODE_IMAGE}
    docker rmi -f $(docker images --filter "dangling=true" -q)
}

# Proxy
export WEBSERVER_IMAGE=go-networking/webserver:current
export HTTPUTIL_PROXY_IMAGE=go-networking/httputilproxy:current
export CUSTOM_PROXY_IMAGE=go-networking/customproxy:current
function build_proxy(){
    docker-compose -f ./build/builder.yml build webserver
    docker-compose -f ./build/builder.yml build httputilproxy
    docker-compose -f ./build/builder.yml build customproxy
}

function clean_proxy(){
    docker rmi -f ${WEBSERVER_IMAGE}
    docker rmi -f ${HTTPUTIL_PROXY_IMAGE}
    docker rmi -f ${CUSTOM_PROXY_IMAGE}
    docker rmi -f $(docker images --filter "dangling=true" -q)
}
