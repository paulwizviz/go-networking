# Peer-to-Peer

In this section we'll demonstrate techniques to implement peer-to-peer networking architecture based principally on this library [Libp2p-go](https://github.com/libp2p/go-libp2p).

## Multi Addressing

The ability of a peer to identify its own IP address is a key element of a peer-to-peer system, since this there is no central entity to managed ip addresses.

* [Source code](../cmd/p2p/multiaddr) demonstrating the concept of multiaddress format.
* Use this [runtime](../deployments/playground.yaml) to play with this example 

## Pinging Nodes

In this example, we demonstrate operations where a running node receives a ping from another node from start state.

The working example is here [./deployments/p2p-ex1.yaml](../deployments/p2p-ex1.yaml)

To see the example in action, run the following commands:

* `./scripts/p2p.sh image build` to create images
* `./scripts/p2p.sh ex1 node1:start` to activate `node1`
* Copy the address of node1 and update the `-laddr` flag value with one from `node1`
* `./scripts/p2p.sh ex1 node2:start` to activate `node2`

## Discovery service using Multicast DNS

This example uses multicast DNS (mdns) as a basis for service discovery. This will only work in local network.

* [Node source code](../cmd/p2p/discovery/main.go)
* [Deployment](../deployments/p2p-ex2.yaml)
* `./scripts/p2p.sh ex2 [start|stop]` -- use this script to run the network

## Discovery service using KHDT

This example uses is a replication of the official example [https://github.com/libp2p/go-libp2p/blob/master/examples/chat-with-rendezvous](https://github.com/libp2p/go-libp2p/blob/master/examples/chat-with-rendezvous). The example here is use docker.

* [Node source code](../cmd/p2p/routing/main.go)
* [Deployment](../deployments/p2p-ex3.yaml)
* `scripts/p2p/sh ex3 [node1:start | node2:start]`

## References

* Libp2p 
    * [Libp2p official Go documentation](https://docs.libp2p.io/guides/getting-started/go)
    * [Getting started with libp2p in Go](https://dev.to/feliperosa/getting-started-with-libp2p-in-go-4hoa)
    * [Building an echo application with libp2p](https://ldej.nl/post/building-an-echo-application-with-libp2p/)