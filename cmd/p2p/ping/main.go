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

	// Get IP address from network
	netAddr, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}

	// Prepare multiaddress format from network and port
	maddr := fmt.Sprintf("/ip4/%s/tcp/%d", strings.Split(netAddr[1].String(), "/")[0], *port)

	// instantiate a peer node
	node, err := libp2p.New(
		libp2p.ListenAddrStrings(maddr),
		libp2p.Ping(false),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer node.Close()

	// Instantiate pinging service
	ps := &ping.PingService{Host: node}
	node.SetStreamHandler(ping.ID, ps.PingHandler)

	// Instantiate a peer info
	peerInfo := peer.AddrInfo{
		ID:    node.ID(),
		Addrs: node.Addrs(),
	}

	// Get Multi Addresses from peer info
	addrs, err := peer.AddrInfoToP2pAddrs(&peerInfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Listerner address: %v\n", addrs[0])

	// Local address of listener peer
	if *laddr != "" {

		// Instantiate multi address from listener address
		addr, err := multiaddr.NewMultiaddr(*laddr)
		if err != nil {
			log.Fatal(err)
		}

		// Instantiate peer info
		p, err := peer.AddrInfoFromP2pAddr(addr)
		if err != nil {
			log.Fatal(err)
		}

		// Connect current node to listener
		err = node.Connect(context.Background(), *p)
		if err != nil {
			log.Fatal("Connect", err)
		}

		// Send ping to listener
		fmt.Println("Sending 5 ping to ", addr)
		ch := ps.Ping(context.Background(), p.ID)
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
