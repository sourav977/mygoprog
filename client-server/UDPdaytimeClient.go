package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkErrorc(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkErrorc(err)
	fmt.Println("ok1")
	_, err = conn.Write([]byte("anything"))
	checkErrorc(err)
	var buf [512]byte
	n, addr, err := conn.ReadFromUDP(buf[0:])
	fmt.Println(addr.IP, addr.Port)
	checkErrorc(err)
	fmt.Println(string(buf[0:n]))
	os.Exit(0)
}
func checkErrorc(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}
