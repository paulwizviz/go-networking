#!/bin/bash

if [ "$(basename $(realpath .))" != "go-networking" ]; then
    echo "You are outside of the project"
    exit 0
else
    . ./scripts/images.sh
    . ./scripts/networks.sh
fi

COMMAND="$1"
SUBCOMMAND1="$2"
SUBCOMMAND2="$3"

function image(){
    local cmd="$1"
    case $cmd in
        "build:playground")
            build_playground
            ;;
        "build:proxy")
            build_proxy
            ;;
        "build:p2p")
            build_p2p
            ;;
        "build")
            build_image
            ;;
        "clean:playground")
            clean_playground
            ;;
        "clean:proxy")
            clean_proxy
            ;;
        "clean:p2p")
            clean_p2p
            ;;
        "clean")
            clean_image
            ;;
        *)
            echo "image [command]
            
command:
    build:playground   playground image
    build:proxy        proxy image
    build:p2p          p2p image
    build              all images
    clean:playground   playground image
    clean:proxy        proxy image
    clean:p2p          p2p image
    clean              all images"
            ;;
    esac
}

function network(){
    local cmd=$1
    case $cmd in
        "playground")
            ground_network
            ;;
        "proxy")
            proxy_network $SUBCOMMAND2
            ;;
        "p2p")
            p2p_network $SUBCOMMAND2
            ;;
        "clean")
            clean_network
            clean_containers
            remove_volumes
            ;;
        *)
            echo "Usage: $0 network [type]
type:
    clean       network assets including volumes
    playground  access ubuntu shell playground
    proxy       network demonstrating proxy servers
    p2p         network demonstrating peer-to-peer architecture"
            ;;
    esac
}

case $COMMAND in
    "image")
        image $SUBCOMMAND1
        ;;
    "network")
        network $SUBCOMMAND1
        ;;
    *)
        echo "$0 <command>
commands:
    image     build or clean
    network   clean, start and stop
"
        ;;
esac