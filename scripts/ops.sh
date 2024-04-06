#!/bin/bash

if [ "$(basename $(realpath .))" != "go-networking" ]; then
    echo "You are outside of the project"
    exit 0
elif [ "$(basename $(realpath .))" == "scripts" ]; then
   . ./image.sh
   . ./network.sh
else
    . ./scripts/image.sh
    . ./scripts/network.sh
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
        "proxy")
            proxy_network $NETWORK_OPS
            ;;
        "p2p")
            p2p_network $NETWORK_OPS
            ;;
        "clean")
            clean_network
            ;;
        *)
            echo "Usage: $0 network [type]
type:
    proxy   network example
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