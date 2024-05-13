#!/bin/bash

export WEBSERVER_IMAGE=go-networking/webserver:current
export HTTPUTIL_PROXY_IMAGE=go-networking/httputilproxy:current
export CUSTOM_PROXY_IMAGE=go-networking/customproxy:current
export P2P_NODE_IMAGE=go-networking/p2pnode:current
export PLAYGROUND_IMAGE=go-networking/playground:current

function build_playground(){
    docker-compose -f ./build/builder.yml build playground
}

function build_p2pnode(){
    docker-compose -f ./build/builder.yml build p2pnode
}

function build_proxy(){
    docker-compose -f ./build/builder.yml build webserver
    docker-compose -f ./build/builder.yml build httputilproxy
    docker-compose -f ./build/builder.yml build customproxy
}

function build_image(){
    docker-compose -f ./build/builder.yml build
}

function clean_p2pnode(){
    docker rmi -f ${P2P_NODE_IMAGE}
    docker rmi -f $(docker images --filter "dangling=true" -q)
}

function clean_playground(){
    docker rmi -f ${PLAYGROUND_IMAGE}
}

function clean_proxy(){
    docker rmi -f ${WEBSERVER_IMAGE}
    docker rmi -f ${HTTPUTIL_PROXY_IMAGE}
    docker rmi -f ${CUSTOM_PROXY_IMAGE}
}


function clean_image(){
    clean_playground
    clean_p2p
    clean_proxy
    docker rmi -f $(docker images --filter "dangling=true" -q)
}

function image(){
    local cmd="$1"
    case $cmd in
        "build:playground")
            build_playground
            ;;
        "build:proxy")
            build_proxy
            ;;
        "build:p2pnode")
            build_p2pnode
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
        "clean:p2pnode")
            clean_p2pnode
            ;;
        "clean")
            clean_image
            ;;
        *)
            echo "image [command]
            
command:
    build:playground   playground image
    build:proxy        proxy image
    build:p2pnode      p2p node image
    build              all images
    clean:playground   playground image
    clean:proxy        proxy image
    clean:p2pnode      p2p node image
    clean              all images"
            ;;
    esac
}