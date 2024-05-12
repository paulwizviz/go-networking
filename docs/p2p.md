# Peer-to-Peer

Peer-to-peer (P2P) network refers to a collection of computing platforms -- i.e. peers -- where every peer consumes and supply resources from other peers.

Throughout this project, we'll be using two techniques when discussing p2p development:

* Standard library networking packages
* Libp2p-go

## Self Addressing

The ability of a peer to identify its own IP address is a key element of a peer-to-peer system, since this there is no central entity to managed ip addresses.

Please refer to this [working example](../cmd/ping/selfaddr/main.go) demonstrating self addressing using standard library and libp2p.
   
## Useful References

* Standard library
    * [Get Local Network Address Using Golang — A Beginner’s Guide](https://systemweakness.com/get-local-network-address-using-golang-a-beginners-guide-7e4074287a03)

* Libp2p 
    * [Libp2p official documentation](https://docs.libp2p.io/guides/getting-started/go)
    * [libp2p-pubsub Peer Discovery with Kademlia DHT](https://medium.com/rahasak/libp2p-pubsub-peer-discovery-with-kademlia-dht-c8b131550ac7)
    * [Getting started with libp2p in Go](https://dev.to/feliperosa/getting-started-with-libp2p-in-go-4hoa)

* Gossip Protocol
    * [Parallel & Distributed Computing - Gossip Protocol](https://www.youtube.com/watch?v=qJpPjzg44R8)