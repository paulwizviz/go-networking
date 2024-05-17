# Network programming

This section discuss techniques for network progamming at level 1 to 4 of the OSI model.

* [Self Addressing](#self-addressing)

## Self Addressing

Here we use the standard library to obtain the network address.

* Source code
    * [Standard library](../cmd/stdlib/selfaddr/main.go) - Using the standard package
    * [Libp2p](../cmd/p2p/multiaddr/main.go) - Using libp2p
* Use this [runtime](../deployments/playground.yaml) to test the working example.

## UDP Transport Programming

In this example, we demonstrate a simple client-server UDP transport protocol.

* [Source code](../cmd/stdlib/transport/main.go) -- Using standard package.
* [Deployment](../deployments/transport.yaml) -- docker-compose script
* [Scripts](../scripts/transport.sh) -- scripts to start/stop network

To use the script:

* Build the docker image: `./scripts/transport.sh image build`
* Start network: `./scripts/transport.sh ops start`
* Shell into a client node: `./scripts/transport.sh shell`
    * In the shell run the app `transport -listener=false -proto=udp -address=192.168.0.2 -port=3030 -msg="Hello"`
* Stop the network: `./scripts/transport.sh ops stop` or `./scripts/transport.sh clean`

## References

* [Get Local Network Address Using Golang — A Beginner’s Guide](https://systemweakness.com/get-local-network-address-using-golang-a-beginners-guide-7e4074287a03)