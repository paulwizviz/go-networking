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

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/protocol/ping"
	multiaddr "github.com/multiformats/go-multiaddr"
)

func main() {

	port := flag.Int("port", 2020, "Port number")
	laddr := flag.String("laddr", "", "Listener address")
	flag.Parse()

	ad, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	maddr := fmt.Sprintf("/ip4/%s/tcp/%d", strings.Split(ad[1].String(), "/")[0], *port)

	node, err := libp2p.New(
		libp2p.ListenAddrStrings(maddr),
		libp2p.Ping(false),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer node.Close()

	ps := &ping.PingService{Host: node}
	node.SetStreamHandler(ping.ID, ps.PingHandler)

	pi := peer.AddrInfo{
		ID:    node.ID(),
		Addrs: node.Addrs(),
	}
	addrs, err := peer.AddrInfoToP2pAddrs(&pi)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Listerner address: %v\n", addrs[0])

	if *laddr != "" {
		addr, err := multiaddr.NewMultiaddr(*laddr)
		if err != nil {
			log.Fatal(err)
		}
		peer, err := peer.AddrInfoFromP2pAddr(addr)
		if err != nil {
			log.Fatal(err)
		}

		err = node.Connect(context.Background(), *peer)
		if err != nil {
			log.Fatal("Connect", err)
		}
		fmt.Println("Sending 5 ping to ", addr)
		ch := ps.Ping(context.Background(), peer.ID)
		for i := 0; i < 5; i++ {
			res := <-ch
			fmt.Println("pinged", addr, "in", res.RTT)
		}
	} else {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch
		fmt.Println("Received signal, shutting down...")
	}
}
