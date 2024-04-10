#!/bin/bash

export P2P_VOLUME="go-network_p2p"
export SOCKET_VOL="go-network_socket"

function create_volume(){
    docker volume create ${P2P_VOLUME}
}

function remove_volume(){
    docker volume rm ${P2P_VOLUME}
}