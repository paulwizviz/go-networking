package main

import (
	"fmt"
	"net"
)

func main() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok {
			if ipNet.IP.IsLoopback() {
				fmt.Println("Loopback IP Address: ", ipNet.IP)
			}
			if ipNet.IP.To4() != nil {
				fmt.Println("IP4 Address", ipNet.IP)
			}
		}
	}
}
