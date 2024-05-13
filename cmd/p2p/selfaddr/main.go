package main

import (
	"fmt"
	"net"

	"github.com/libp2p/go-libp2p"
)

func libp2pEx() {
	// Start a libp2p address
	node, err := libp2p.New()
	if err != nil {
		fmt.Printf("New node error: %v", err)
	}
	defer node.Close()

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
}

func stdlib() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok {
			if ipNet.IP.IsLoopback() {
				fmt.Println("Loopback IP Address: ", ipNet.IP)
			}
			if ipNet.IP.To4() != nil {
				fmt.Println("IP4 Address", ipNet.IP)
			}
		}
	}
}

func main() {
	fmt.Println("-- Libp2p --")
	libp2pEx()
	fmt.Println("-- Standard library --")
	stdlib()
}
