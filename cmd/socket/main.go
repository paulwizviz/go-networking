package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

// client operation dials a connection
func client(socket string, msg string) error {
	conn, err := net.Dial("unix", socket)
	if err != nil {
		return err
	}
	_, err = conn.Write([]byte(msg))
	if err != nil {
		return err
	}
	return nil
}

func processConn(c net.Conn) {
	buf := make([]byte, 1024)
	nr, err := c.Read(buf)
	if err != nil {
		log.Fatalf("Unable to read. Reason: %v", err)
	}
	fmt.Printf("Received: %s\n", string(buf[:nr]))
}

func server(socket string) error {

	// socket listening
	l, err := net.Listen("unix", socket)
	if err != nil {
		log.Fatalf("Unable to start server. Reason: %s", err.Error())
	}
	defer os.Remove(socket)

	for {
		fd, err := l.Accept()
		if err != nil {
			log.Fatalf("Error: %s", err.Error())
		}
		processConn(fd)
	}
}

var isServer bool
var socketAddr string

func main() {

	flag.StringVar(&socketAddr, "socket", "/var/app/test.socket", "socket address")
	flag.BoolVar(&isServer, "server", true, "server flag")
	flag.Parse()

	if isServer {
		err := server(socketAddr)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		client(socketAddr, "Hello")
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		os.Exit(0)
	}
}
