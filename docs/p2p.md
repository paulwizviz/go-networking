# Peer-to-Peer

Peer-to-peer (P2P) network refers to a collection of computing platforms -- i.e. peers -- where every peer consumes and supply resources from other peers.

There are two key concepts in P2P network:

* Peer discovery - This is the process of a peer finding and announcing services to other peers.
* Peer routing - This is a process of finding a specific peer.

Throughout this project, we'll be using two techniques when discussing p2p development:

* Standard library networking packages
* Libp2p-go

## Self Addressing

The ability of a peer to identify its own IP address is a key element of a peer-to-peer system, since this there is no central entity to managed ip addresses.

Please refer to this [working example](../cmd/ping/selfaddr/main.go) demonstrating self addressing using standard library and libp2p.
   
## Useful References

* [Libp2p official documentation](https://docs.libp2p.io/guides/getting-started/go)
* [Get Local Network Address Using Golang — A Beginner’s Guide](https://systemweakness.com/get-local-network-address-using-golang-a-beginners-guide-7e4074287a03)