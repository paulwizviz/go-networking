package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/discovery"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery/routing"
	"github.com/multiformats/go-multiaddr"
)

type RoutingTable struct {
	h         host.Host
	providers map[string]map[peer.ID]peer.AddrInfo
}

func (r *RoutingTable) Provide(ctx context.Context, cid cid.Cid, bcast bool) error {
	pmap, ok := r.providers[cid.String()]
	if !ok {
		pmap = make(map[peer.ID]peer.AddrInfo)
		r.providers[cid.String()] = pmap
	}
	pmap[r.h.ID()] = peer.AddrInfo{ID: r.h.ID(), Addrs: r.h.Addrs()}
	return nil
}

func (r *RoutingTable) FindProvidersAsync(ctx context.Context, cid cid.Cid, limit int) <-chan peer.AddrInfo {
	ch := make(chan peer.AddrInfo)
	go func() {
		defer close(ch)

		pmap, ok := r.providers[cid.String()]
		if !ok {
			return
		}

		for _, pi := range pmap {
			select {
			case ch <- pi:
			case <-ctx.Done():
				return
			}
		}
	}()
	return ch
}

func main() {

	bootAddr := flag.String("boot-addr", "", "bootstrap address")
	flag.Parse()

	if *bootAddr == "" {
		log.Fatal("No boot strap address")
	}

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

	bootstrapAddr, err := multiaddr.NewMultiaddr(*bootAddr)
	if err != nil {
		log.Fatal(err)
	}
	bootstrapPeer, err := peer.AddrInfoFromP2pAddr(bootstrapAddr)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	if err := host.Connect(ctx, *bootstrapPeer); err != nil {
		log.Printf("Failed to connect to bootstrap node: %v", err)
	} else {
		log.Printf("Connected to bootstrap node: %s", bootstrapPeer.ID.String())
	}

	kDHT, err := dht.New(ctx, host, dht.Mode(dht.ModeAuto))
	if err != nil {
		log.Fatal(err)
	}

	// Set up routing discovery using the DHT
	rd := routing.NewRoutingDiscovery(kDHT)

	// Advertise our presence
	_, err = rd.Advertise(ctx, "my-libp2p-app", discovery.TTL(1*time.Second))
	if err != nil {
		log.Fatal("Advertise Failed: ", err)
	}

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

			fmt.Println("Discovered peers:")
			for p := range peers {
				if p.ID != host.ID() {
					fmt.Println(p.ID.String())
				}
			}
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("Received signal, shutting down...")
}
