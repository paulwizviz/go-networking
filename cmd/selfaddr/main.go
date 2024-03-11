// This example shows two approaches to obtaining interfaces IP address, using:
// * standard library.
// * go-libp2p.
package main

import (
	"fmt"
	"net"

	"github.com/libp2p/go-libp2p"
)

// stdlib demonstrate a technique to obtain network interface address.
// NOTE: This does not spin up a running node.
func stdlib() {
	fmt.Println("--- Using standard library ---")
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			fmt.Println("IP Address: ", ipNet.IP)
		}
	}
}

// golibp2p this example uses go-libp2p to obtain network interfaces
// and starting a node.
func golibp2p() {
	fmt.Println("--- Using libp2p ---")

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

func main() {
	stdlib()
	golibp2p()
}
