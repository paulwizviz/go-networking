#!/bin/bash

export P2P_VOLUME="go-network-vol"

function create_volume(){
    docker volume create ${P2P_VOLUME}
}

function remove_volume(){
    docker volume rm ${P2P_VOLUME}
}