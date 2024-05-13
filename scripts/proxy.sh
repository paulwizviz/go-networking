#!/bin/sh

if [ "$(basename $(realpath .))" != "go-networking" ]; then
    echo "You are outside of the project"
    exit 0
else
    . ./scripts/images.sh
fi

COMMAND=$1
SUBCOMMAND=$2

function image(){
    local cmd=$1
    case $cmd in
        "build")
            build_proxy
            ;;
        "clean")
            clean_proxy
            ;;
        *)
            ;;
    esac
}

export PROXY_NETWORK="go-networking_proxy"
function ops(){
    local cmd=$1
    case $cmd in
        "start")
            docker-compose -f ./deployments/proxy-network.yaml up
            ;;
        "stop")
            docker-compose -f ./deployments/proxy-network.yaml down
            ;;
        *)
            echo "Usage: $0 ops [command]
            
command:
    start   proxy network
    stop    proxy network"
            ;;
    esac
}


case $COMMAND in
    "image")
        image $SUBCOMMAND
        ;;
    "ops")
        ops $SUBCOMMAND
        ;;
    *)
        echo "Usage: $0 [command]

command:
    image    related operations
    ops      network operations"
        ;;
esac
