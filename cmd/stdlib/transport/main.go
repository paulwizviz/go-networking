package main

import (
	"flag"
	"log"
	"net"
)

func udpListener(proto string, laddr string, port int) error {

	udpAddr := &net.UDPAddr{
		IP:   net.ParseIP(laddr),
		Port: port,
	}

	conn, err := net.ListenUDP(proto, udpAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	p := make([]byte, 256)
	for {
		_, err := conn.Read(p)
		if err != nil {
			log.Println(err)
		}
		log.Printf("String: %s Bytes: %v", string(p), p)
		if string(p[:4]) == "stop" {
			log.Println("==break==")
			break
		}
	}
	return nil
}

func udpClient(proto string, addr string, port int, msg string) error {
	udpAddr := &net.UDPAddr{
		IP:   net.ParseIP(addr),
		Port: port,
	}
	conn, err := net.DialUDP(proto, nil, udpAddr)
	if err != nil {
		return err
	}
	_, err = conn.Write([]byte(msg))
	if err != nil {
		return err
	}
	return nil
}

func main() {

	isListener := flag.Bool("listener", true, "Is a listener")
	proto := flag.String("proto", "", "Network protocol, udp or tcp")
	addr := flag.String("address", "", "IPv4 address")
	port := flag.Int("port", 3030, "Port")
	msg := flag.String("msg", "", "message")
	flag.Parse()

	log.Printf("IsListener: %v Proto: %s Address: %s Port: %d", *isListener, *proto, *addr, *port)

	if *isListener && *proto == "udp" {
		log.Println("listener started")
		err := udpListener(*proto, *addr, *port)
		if err != nil {
			log.Println(err)
		}
	}

	if !*isListener && *proto == "udp" {
		log.Println("client started")
		err := udpClient(*proto, *addr, *port, *msg)
		if err != nil {
			log.Println(err)
		}
	}
}
