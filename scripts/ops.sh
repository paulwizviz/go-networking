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

case $COMMAND in
    "clean")
        clean_network
        clean_containers
        remove_volumes
        ;;
    "image")
        image $SUBCOMMAND1
        ;;
    "playground")
        ground_network
        ;;
    "proxy")
        proxy_network $SUBCOMMAND1
        ;;
    "p2p")
        p2p_network $SUBCOMMAND1
        ;;
    *)
        echo "$0 <command>
commands:
    clean       project artefacts
    image       build or clean
    playground  access ubuntu shell playground
    proxy       network demonstrating proxy servers
    p2p         network demonstrating peer-to-peer architecture
"
        ;;
esac