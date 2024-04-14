#!/bin/bash

if [ "$(basename $(realpath .))" != "go-networking" ]; then
    echo "You are outside of the project"
    exit 0
else
    . ./scripts/images.sh
    . ./scripts/networks.sh
fi

COMMAND="$1"
SUBCOMMAND="$2"
NETWORK_OPS="$3"

function image(){
    local cmd="$1"
    case $cmd in
        "build")
            build_image
            ;;
        "clean")
            clean_image
            ;;
        *)
            echo "image [ build | clean]"
            ;;
    esac
}

function network(){
    local cmd=$1
    case $cmd in
        "playground")
            playground
            ;;
        "proxy")
            proxy_network $NETWORK_OPS
            ;;
        "p2p")
            p2p_network $NETWORK_OPS
            ;;
        "clean")
            clean_network
            remove_containers
            remove_volumes
            ;;
        *)
            echo "Usage: $0 network [type]
type:
    clean       network assets including volumes
    playground  access ubuntu shell playground
    proxy       network demonstrating proxy servers
    p2p         network demonstrating peer-to-peer architecture
"
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