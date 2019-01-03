package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":1200"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)
	fmt.Println(udpAddr.IP, udpAddr.Port)
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)
	for {
		handleClient(conn)
	}
}
func handleClient(conn *net.UDPConn) {
	var buf [512]byte
	_, addr, err := conn.ReadFromUDP(buf[0:])
	fmt.Println(addr.IP, addr.Port)
	if err != nil {
		return
	}
	daytime := time.Now().String()
	_, errr := conn.WriteToUDP([]byte(daytime), addr)
	fmt.Println(addr.IP, addr.Port)
	checkError(errr)
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}
