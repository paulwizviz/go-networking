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
            build_transport
            ;;
        "clean")
            clean_transport
            ;;
        *)
            echo "Usage: $0 image [command]
            
command:
    build    transport image
    clean    transport image"
            ;;
    esac
}

export TRANSPORT_NETWORK="go-networking_transport"
function ops(){
    local cmd=$1
    case $cmd in
        "start")
            docker-compose -f ./deployments/transport.yaml up
            ;;
        "shell")
            docker-compose -f ./deployments/transport.yaml run --rm -it client /bin/sh
            ;;
        "stop")
            docker-compose -f ./deployments/transport.yaml down
            ;;
        *)
            echo "Usage: $0 ops [command]
            
command:
    start  network
    stop   network"
        ;;
    esac
}

case $COMMAND in
    "clean")
        ops stop
        image clean
        docker network rm $TRANSPORT_NETWORK
        ;;
    "images")
        image $SUBCOMMAND
        ;;
    "ops")
        ops $SUBCOMMAND
        ;;
    *)
        echo "Usage: $0 [command]
command:
    clean   project artefacts
    images  related operations
    ops     network operations"
        ;;
esac