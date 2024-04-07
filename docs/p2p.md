# Peer-to-Peer

Peer-to-peer (P2P) network refers to a collection of computing platforms -- i.e. peers -- where every peer consumes and supply resources from other peers.

There are two key concepts in P2P network:

* Peer discovery - This is the process of a peer finding and announcing services to other peers.
* Peer routing - This is a process of finding a specific peer.

## Using standard library

This section discuss the use of standard library to build network applications.

### Working examples

* [Self Addressing](../cmd/netpkg/selfaddr/main.go) - [source](https://systemweakness.com/get-local-network-address-using-golang-a-beginners-guide-7e4074287a03)
    * This example demonstrate the use of standard library networking to extract IP address from your platform

## Using Libp2p

Please refer to [official documentation](https://docs.libp2p.io/guides/getting-started/go)

The following are examples of applications based on libp2p:

* [Web3 basic](https://pl-launchpad.io/curriculum/web3/objectives/)
* [IPFS basic](https://pl-launchpad.io/curriculum/ipfs/objectives/)

### Working examples

* [Self addressing](../cmd/selfaddr/main.go)
    * This example demonstrate the use of libp2p-go to extract IP address from your platform.