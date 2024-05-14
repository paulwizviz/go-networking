# Peer-to-Peer

In this section we'll demonstrate techniques to implement peer-to-peer networking architecture. We'll be using these components to help us build examples of p2p.

* Standard library networking packages
* [Libp2p-go](https://github.com/libp2p/go-libp2p)

## Self Addressing

The ability of a peer to identify its own IP address is a key element of a peer-to-peer system, since this there is no central entity to managed ip addresses.

Please refer to this [working example](../cmd/ping/selfaddr) demonstrating self addressing using standard library and libp2p.

Use this together with the [playground container](../deployments/playground.yaml) 

## Pinging Nodes

In this example, we demonstrate operations where a running node receives a ping from another node from start state.

The working example is here [./deployments/p2p-ex1.yaml](../deployments/p2p-ex1.yaml)

To see the example in action, run the following commands:

* `./scripts/p2p.sh image build` to create images
* `./scripts/p2p.sh ex1 node1:start` to activate `node1`
* Copy the address of node1 and update the `-laddr` flag value with one from `node1`
* `./scripts/p2p.sh ex1 node2:start` to activate `node2`

## References

* Standard library
    * [Get Local Network Address Using Golang — A Beginner’s Guide](https://systemweakness.com/get-local-network-address-using-golang-a-beginners-guide-7e4074287a03)

* Libp2p 
    * [Libp2p official documentation](https://docs.libp2p.io/guides/getting-started/go)
    * [libp2p-pubsub Peer Discovery with Kademlia DHT](https://medium.com/rahasak/libp2p-pubsub-peer-discovery-with-kademlia-dht-c8b131550ac7)
    * [Getting started with libp2p in Go](https://dev.to/feliperosa/getting-started-with-libp2p-in-go-4hoa)
  