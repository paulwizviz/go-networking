#!/bin/sh

if [ "$(basename $(realpath .))" != "go-networking" ]; then
    echo "You are outside of the project"
    exit 0
else
    . ./scripts/images.sh
fi

export P2P_NETWORK_EX1="go-networking_p2p-ex1-net"
export P2P_VOLUME_EX1="go-networking_p2p-ex1-vol"

COMMAND=$1
SUBCOMMAND1=$2

function image(){
    local cmd=$1
    case $cmd in
        "build")
            build_p2pnode
            ;;
        "clean")
            clean_p2pnode
            ;;
    esac
}

function ex1(){
    local cmd=$1
    case $cmd in
        "node1:start")
            docker-compose -f ./deployments/p2p-ex1.yaml up node1
            ;;
        "node1:stop")
            docker-compose -f ./deployments/p2p-ex1.yaml down node1
            ;; 
        "node2:start")
            docker-compose -f ./deployments/p2p-ex1.yaml up node2
            ;;
        "node2:stop")
            docker-compose -f ./deployments/p2p-ex1.yaml down node2
            ;;
        "stop")
            docker-compose -f ./deployments/p2p-ex1.yaml down
            ;;
        *)
            echo "Usage: $0 ex1 [command]
            
command:
    node1:start    activate node1
    node1:stop     deactivate node1
    node2:start    activate node2
    node2:stop     deactivate node2
    stop           all nodes"
            ;;
    esac
}

case $COMMAND in
    "image")
        image $SUBCOMMAND1
        ;;
    "ex1")
        ex1 $SUBCOMMAND1
        ;;
    *)
        echo "Usage: $0 [command]
        
command:
    ex1   operations for example 1"
        ;;
esac