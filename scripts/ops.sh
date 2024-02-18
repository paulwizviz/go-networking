#!/bin/bash

export WEBSERVER_IMAGE=go-networking/webserver:current
export HTTPUTIL_PROXY_IMAGE=go-networking/httputilproxy:current
export CUSTOM_PROXY_IMAGE=go-networking/customproxy:current
export NETWORK=go-networking_network

COMMAND="$1"
SUBCOMMAND="$2"

function image(){
    local cmd="$1"
    case $cmd in
        "build")
            docker-compose -f ./build/builder.yml build
            ;;
        "clean")
            docker rmi -f ${WEBSERVER_IMAGE}
            docker rmi -f ${HTTPUTIL_PROXY_IMAGE}
            docker rmi -f ${CUSTOM_PROXY_IMAGE}
            docker rmi -f $(docker images --filter "dangling=true" -q)
            ;;
        *)
            echo "image [ build | clean]"
            ;;
    esac
}

function network(){
    local cmd=$1
    case $cmd in
        "clean")
            docker-compose -f ./deployment/docker-compose.yml down
            rm -rf ./tmp
            ;;
        "start")
            docker-compose -f ./deployment/docker-compose.yml up
            ;;
        "stop")
            docker-compose -f ./deployment/docker-compose.yml down
            ;;
        *)
            echo "network [ clean | start | stop ]"
            ;;
    esac
}

case $COMMAND in
    "image")
        image $SUBCOMMAND
        ;;
    "network")
        network $SUBCOMMAND
        ;;
    *)
        echo "$0 <command>
commands:
    image     build or clean
    network   clean, start and stop
"
        ;;
esac