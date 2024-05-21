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

	"github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
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

	ctx := context.Background()

	// Set up the DHT
	kDHT, err := dht.New(ctx, host, dht.Mode(dht.ModeAuto))
	if err != nil {
		log.Fatal(err)
	}

	// Bootstrap the DHT
	err = kDHT.Bootstrap(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Print this node's addresses and ID
	fmt.Println("Addresses:", host.Addrs())
	fmt.Println("ID:", host.ID())

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("Received signal, shutting down...")
}
