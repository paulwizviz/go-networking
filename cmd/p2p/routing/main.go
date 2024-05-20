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

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery/routing"
)

type RoutingTable struct {
}

func (r *RoutingTable) Provide(ctx context.Context, cid cid.Cid, bcast bool) error {
	log.Println(cid.String(), bcast)
	return nil
}

func (r *RoutingTable) FindProvidersAsync(ctx context.Context, cid cid.Cid, limit int) <-chan peer.AddrInfo {

	return nil
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

	rd := routing.NewRoutingDiscovery(&RoutingTable{})
	rd.Advertise(context.TODO(), "test-app")
	peerInfo, err := rd.FindPeers(context.TODO(), "test-app")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for pi := range peerInfo {
			if pi.ID == host.ID() {
				continue // Skip self
			}
			fmt.Printf("Found peer: %s\n", pi.ID.String())

			// Attempt to connect to the discovered peer
			if err := host.Connect(context.TODO(), pi); err != nil {
				fmt.Printf("Failed to connect to peer %s: %s\n", pi.ID.String(), err)
			} else {
				fmt.Printf("Connected to peer %s\n", pi.ID.String())
			}
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("Received signal, shutting down...")
}
