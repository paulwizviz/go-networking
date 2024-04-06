#!/bin/bash

export NETWORK=go-networking_network

function proxy_network(){
    local cmd=$1
    case $cmd in
        "start")
            docker-compose -f ./deployment/proxy-network.yaml up
            ;;
        "stop")
            docker-compose -f ./deployment/proxy-network.yaml down
            ;;
        *)
            echo "Usage: $0 network proxy [command]
command:
    start    network
    stop     network
"
            ;;
    esac
}

function p2p_network(){
    local cmd=$1
    case $cmd in
        "start")
            docker-compose -f ./deployment/p2p-network.yaml up
            ;;
        "stop")
            docker-compose -f ./deployment/p2p-network.yaml down
            ;;
        *)
            echo "Usage: $0 network p2p [command]
command:
    start    network
    stop     network
"
            ;;
    esac
}

function clean_network(){
    docker-compose -f ./deployment/proxy-network.yaml down
    docker-compose -f ./deployment/p2p-network.yaml down
    docker network rm $NETWORK
}