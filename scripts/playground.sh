#!/bin/sh

if [ "$(basename $(realpath .))" != "go-networking" ]; then
    echo "You are outside of the project"
    exit 0
else
    . ./scripts/images.sh
fi

COMMAND=$1
SUBCOMMAND=$2

export PLAYGROUND_NETWORK="go-networking_playground-net"
export PLAYGROUND_CONTAINER="playground"

function image(){
    local cmd=$1
    case $cmd in
        "build")
            build_playground
            ;; 
        "clean")
            clean_playground 
            ;;
        *)
            echo "Usage: $0 image [command]
            
command:
    build   image
    clean   remove images from cache"
    esac
}

function ops(){
    local cmd=$1
    case $cmd in
        "shell")
            docker-compose -f ./deployments/playground.yaml run --rm -it playground /bin/sh
            ;;
        "clean")
            docker container rm ${PLAYGROUND_CONTAINER}
            docker volume rm ${PLAYGROUND_VOLUME}
            ;;
        *)
            echo "Usage: $0 ops [command]
            
command:
    shell   into playground container
    clean   network artefacts"
            ;;
    esac
}

case $COMMAND in
    "images")
        image $SUBCOMMAND
        ;;
    "ops")
        ops $SUBCOMMAND
        ;;
    *)
        echo "Usage: $0 [command]
        
command:
    images  related operations
    ops     network operations"
        ;;
esac