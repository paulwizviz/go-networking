#!/bin/sh

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
            docker-compose -f ./deployments/p2p-network.yaml up
            ;;
        "stop")
            docker-compose -f ./deployments/p2p-network.yaml down
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

export PLAYGROUND_CONTAINER="playground"
function playground_network(){
    docker-compose -f ./deployments/playground.yaml run --rm -it playground /bin/sh
}

function clean_containers(){
    docker rm -f ${PLAYGROUND_CONTAINER}
}

# Volumes
export P2P_VOLUME="go-network_p2p"
export PLAYGROUND_VOLUME="go-network_playground"

function remove_volumes(){
    docker volume rm ${P2P_VOLUME}
    docker volume rm ${PLAYGROUND_VOLUME}
}

function clean_network(){
    clean_containers
    p2p_network stop
    proxy_network stop
    remove_volumes
    docker network rm $NETWORK
}