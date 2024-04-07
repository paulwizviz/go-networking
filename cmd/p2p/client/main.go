package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/peer"
)

func main() {

	port := flag.Int("port", 2020, "Port number")
	flag.Parse()

	maddr := fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", *port)
	fmt.Println(maddr)

	node, err := libp2p.New(
		libp2p.ListenAddrStrings(maddr),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer node.Close()

	pi := peer.AddrInfo{
		ID:    node.ID(),
		Addrs: node.Addrs(),
	}

	addrs, err := peer.AddrInfoToP2pAddrs(&pi)
	if err != nil {
		log.Fatal(err)
	}

	for i, addr := range addrs {
		fmt.Printf("%d: %v\n", i, addr)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("Received signal, shutting down...")
}
