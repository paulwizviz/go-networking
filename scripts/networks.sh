#!/bin/bash

# Networks
export NETWORK=go-networking_network

function proxy_network(){
    local cmd=$1
    case $cmd in
        "start")
            docker-compose -f ./deployments/proxy-network.yaml up
            ;;
        "stop")
            docker-compose -f ./deployments/proxy-network.yaml down
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

# Containers
export PLAYGROUND_CONTAINER="playground"

function playground(){
    docker-compose -f ./deployments/playground.yaml run --rm -it playground /bin/bash
}

function clean_containers(){
    docker rm -f ${PLAYGROUND_CONTAINER}
}

# Volumes
export P2P_VOLUME="go-network_p2p"
export SOCKET_VOLUME="go-network_socket"
export PLAYGROUND_VOLUME="go-network_playground"

function remove_volumes(){
    docker volume rm ${P2P_VOLUME}
    docker volume rm ${PLAYGROUND_VOLUME}
    docker volume rm ${SOCKET_VOLUME}
}