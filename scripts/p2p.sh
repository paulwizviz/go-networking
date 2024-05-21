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

export P2P_NETWORK_EX1="go-networking_p2p-ex1-net"
export P2P_VOLUME_EX1="go-networking_p2p-ex1-vol"

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
        "clean")
            ex1 stop
            docker network rm ${P2P_NETWORK_EX1}
            docker volume rm ${P2P_VOLUME_EX1}
            ;;
        *)
            echo "Usage: $0 ex1 [command]
            
command:
    clean          remove ex1 artefacts
    node1:start    activate node1
    node1:stop     deactivate node1
    node2:start    activate node2
    node2:stop     deactivate node2
    stop           all nodes"
            ;;
    esac
}

export P2P_NET_EX2="p2p-ex2-net"
export P2P_VOL_EX2="p2p-ex2-vol"
function ex2(){
    local cmd=$1
    case $cmd in
        "start")
            docker-compose -f ./deployments/p2p-ex2.yaml up
            ;;
        "stop")
            docker-compose -f ./deployments/p2p-ex2.yaml down
            ;;
        "clean")
            ex2 stop
            docker network rm ${P2P_NET_EX2}
            docker volume rm ${P2P_VOL_EX2}
            ;;
        *)
            echo "Usage: $0 ex2 [command]
            
command:
    clean  ex2 artefacts
    start  ex2 network
    stop   network"
            ;;
    esac
}

export P2P_NET_EX3="p2p-ex3-net"
export P2P_VOL_EX3="p2p-ex3-vol"
function ex3(){
    local cmd=$1
    case $cmd in
        "boot:start")
            docker-compose -f ./deployments/p2p-ex3.yaml up boot
            ;;
       "node:start")
            docker-compose -f ./deployments/p2p-ex3.yaml up node1 node2
            ;;
        "stop")
            docker-compose -f ./deployments/p2p-ex3.yaml down
            ;;
        "clean")
            ex3 stop
            docker network rm ${P2P_NET_EX3}
            docker volume rm ${P2P_VOL_EX3}
            ;;
        *)
            echo "Usage: $0 ex3 [command]
            
command:
    boot:start   boot node of ex3
    clean        ex3 artefacts
    node:start  nodes ex3 network
    stop         network"
            ;;
    esac
}


case $COMMAND in
    "images")
        image $SUBCOMMAND1
        ;;
    "ex1")
        ex1 $SUBCOMMAND1
        ;;
    "ex2")
        ex2 $SUBCOMMAND1
        ;;
    "ex3")
        ex3 $SUBCOMMAND1
        ;;
    "clean")
        image clean
        ex1 clean
        ex2 clean
        ex3 clean
        ;;
    *)
        echo "Usage: $0 [command]
        
command:
    images  build or clean operations
    ex1    operations for example 1
    ex2    operations for example 2
    ex3    operations for example 3
    clean  project of artefacts"
        ;;
esac