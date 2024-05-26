package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/ipfs/go-log/v2"
	"github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery/routing"
	"github.com/libp2p/go-libp2p/p2p/discovery/util"
)

var logger = log.Logger("my-libp2p-app")

func main() {
	log.SetAllLoggers(log.LevelWarn)
	log.SetLogLevel("my-libp2p-app", "info")

	laddrs, err := net.InterfaceAddrs()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info(laddrs)
	ipaddr := strings.Split(laddrs[1].String(), "/")

	host, err := libp2p.New(
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/%s/tcp/2020", ipaddr[0])),
	)
	if err != nil {
		logger.Fatal(err)
	}
	defer host.Close()

	// Print this node's addresses and ID
	logger.Info("Addresses:", host.Addrs())
	logger.Info("ID:", host.ID())

	ctx := context.Background()

	bootstrapPeers := make([]peer.AddrInfo, len(dht.DefaultBootstrapPeers))
	for i, addr := range dht.DefaultBootstrapPeers {
		peerinfo, _ := peer.AddrInfoFromP2pAddr(addr)
		bootstrapPeers[i] = *peerinfo
	}
	kDHT, err := dht.New(ctx, host, dht.BootstrapPeers(bootstrapPeers...))
	if err != nil {
		logger.Fatal(err)
	}

	// Boot strapping the DHT
	logger.Debug("Bootstrapping the DHT")
	err = kDHT.Bootstrap(ctx)
	if err != nil {
		logger.Fatal(err)
	}

	// Paused to ensure bootstrap start
	time.Sleep(1 * time.Second)

	// Set up routing discovery using the DHT
	logger.Info("Announcing ourselves")
	rd := routing.NewRoutingDiscovery(kDHT)
	// Advertise our presence
	util.Advertise(ctx, rd, "my libp2p app")
	logger.Info("We have Announced")

	logger.Debug("Searching for other peers...")
	peers, err := rd.FindPeers(ctx, "my libp2p app")
	if err != nil {
		logger.Fatal(err)
	}

	for p := range peers {
		if p.ID == host.ID() {
			continue
		}
		logger.Debug("Found peer:", p)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("Received signal, shutting down...")
}
