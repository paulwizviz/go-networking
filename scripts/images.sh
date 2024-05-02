#!/bin/bash

export WEBSERVER_IMAGE=go-networking/webserver:current
export HTTPUTIL_PROXY_IMAGE=go-networking/httputilproxy:current
export CUSTOM_PROXY_IMAGE=go-networking/customproxy:current
export P2P_CLIENT_IMAGE=go-networking/p2pclient:current
export PLAYGROUND_IMAGE=go-networking/playground:current

function build_playground(){
    docker-compose -f ./build/builder.yml build playground
}

function build_p2p(){
    docker-compose -f ./build/builder.yml build p2p
}

function build_proxy(){
    docker-compose -f ./build/builder.yml build webserver
    docker-compose -f ./build/builder.yml build httputilproxy
    docker-compose -f ./build/builder.yml build customproxy
}

function build_image(){
    build_playground
    build_p2p
    build_proxy
}

function clean_p2p(){
    docker rmi -f ${P2P_CLIENT_IMAGE}
}

function clean_playground(){
    docker rmi -f ${PLAYGROUND_IMAGE}
}

function clean_proxy(){
    docker rmi -f ${WEBSERVER_IMAGE}
    docker rmi -f ${HTTPUTIL_PROXY_IMAGE}
    docker rmi -f ${CUSTOM_PROXY_IMAGE}
}


function clean_image(){
    clean_playground
    clean_p2p
    clean_proxy
    docker rmi -f $(docker images --filter "dangling=true" -q)
}