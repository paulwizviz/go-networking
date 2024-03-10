// This example demonstrate code to extract local ip address
// from network (e.g. en0). On unix based system, run this
// cli command:
//
//	"ipconfig getifaddr en0"
//
// here is an example of the return value: 192.168.1.x
package main

import (
	"fmt"

	"github.com/libp2p/go-libp2p"
)

func main() {
	// Start a libp2p node with default settings
	node, err := libp2p.New()
	if err != nil {
		panic(err)
	}

	// Print the node's listening addresses in multiaddr format.
	// Examples:
	// ip4/127.0.0.1/tcp/61218 <-- localhost
	// ....
	// Listen addresses: /ip4/192.168.1.73/tcp/61218 <-- ipv4 based
	// .....
	// Listen addresses: /ip6/::1/tcp/61221 <-- ipv6 based
	for _, addr := range node.Addrs() {
		fmt.Println("Listen addresses:", addr)
	}

	// shut the node down
	if err := node.Close(); err != nil {
		panic(err)
	}
}
