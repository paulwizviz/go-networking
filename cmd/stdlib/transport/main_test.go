package main

import (
	"fmt"
	"net"
)

func Example_netResolveAddr() {
	addr, err := net.ResolveUDPAddr("udp", "192.168.0.1:2000")
	fmt.Println(addr, err)

	addr, err = net.ResolveUDPAddr("udp", "http://192.168.0.1:2000")
	fmt.Println(addr, err)

	addr, err = net.ResolveUDPAddr("udp", "hello")
	fmt.Println(addr, err)

	addr, err = net.ResolveUDPAddr("udp", "hello:2020")
	fmt.Println(addr, err)

	addr, err = net.ResolveUDPAddr("udp", "localhost:2020")
	fmt.Println(addr, err)

	// Output:
	// 192.168.0.1:2000 <nil>
	// <nil> address http://192.168.0.1:2000: too many colons in address
	// <nil> address hello: missing port in address
	// <nil> lookup hello: no such host
	// 127.0.0.1:2020 <nil>
}
