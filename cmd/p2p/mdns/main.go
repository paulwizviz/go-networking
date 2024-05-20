package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
)

type discoveryNotifee struct {
	h host.Host
}

func (n *discoveryNotifee) HandlePeerFound(peerInfo peer.AddrInfo) {
	log.Println("found peer", peerInfo.String())
}

func main() {

	laddrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}
	ipaddr := strings.Split(laddrs[1].String(), "/")

	host, err := libp2p.New(
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/%s/tcp/2002", ipaddr[0])),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer host.Close()

	// Print this node's addresses and ID
	fmt.Println("Addresses:", host.Addrs())
	fmt.Println("ID:", host.ID())

	// Setup peer discovery.
	discoveryService := mdns.NewMdnsService(host, "discovery", &discoveryNotifee{h: host})
	if err != nil {
		log.Fatal(err)
	}
	defer discoveryService.Close()

	err = discoveryService.Start()
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("Received signal, shutting down...")
}
