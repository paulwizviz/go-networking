package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/ipfs/go-log/v2"
	"github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
)

var logger = log.Logger("my-libp2p-app")

func main() {

	log.SetAllLoggers(log.LevelWarn)
	log.SetLogLevel("my-libp2p-app", "info")

	laddrs, err := net.InterfaceAddrs()
	if err != nil {
		logger.Fatal(err)
	}
	ipaddr := strings.Split(laddrs[1].String(), "/")

	host, err := libp2p.New(
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/%s/tcp/2002", ipaddr[0])),
	)
	if err != nil {
		logger.Fatal(err)
	}
	defer host.Close()

	// Print this node's addresses and ID
	fmt.Println("Addresses:", host.Addrs())
	fmt.Println("ID:", host.ID())

	logger.Debug("New DHT")
	kdht, err := dht.New(context.TODO(), host)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Debug("Get into bootstrap state")
	err = kdht.Bootstrap(context.TODO())
	if err != nil {
		logger.Fatal(err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("Received signal, shutting down...")
}
