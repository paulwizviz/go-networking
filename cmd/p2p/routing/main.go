package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery/routing"
	"github.com/libp2p/go-libp2p/p2p/discovery/util"
)

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

	ctx := context.Background()

	bootstrapPeers := make([]peer.AddrInfo, len(dht.DefaultBootstrapPeers))
	for i, addr := range dht.DefaultBootstrapPeers {
		peerinfo, _ := peer.AddrInfoFromP2pAddr(addr)
		bootstrapPeers[i] = *peerinfo
	}

	kDHT, err := dht.New(ctx, host, dht.BootstrapPeers(bootstrapPeers...))
	if err != nil {
		log.Fatal(err)
	}

	// Set up routing discovery using the DHT
	rd := routing.NewRoutingDiscovery(kDHT)

	if err = kDHT.Bootstrap(ctx); err != nil {
		log.Fatal(err)
	}

	// Advertise our presence
	util.Advertise(ctx, rd, "my-libp2p-app")

	// Discover peers periodically
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			peers, err := rd.FindPeers(ctx, "my-libp2p-app")
			if err != nil {
				log.Printf("Failed to find peers: %v", err)
				continue
			}

			fmt.Printf("Host Address: %v ID: %v\n", host.Addrs(), host.ID())
			for p := range peers {
				if len(p.Addrs) != 0 {
					fmt.Printf("Peer Address: %v ID: %v\n", p.Addrs, p.ID.String())
				}
			}
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("Received signal, shutting down...")
}
