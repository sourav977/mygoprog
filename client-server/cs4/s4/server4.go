package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// only needed below for sample processing

func main() {

	fmt.Println("Launching server...")
	// listen on all interfaces
	service := "127.0.0.1:1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	ln, err1 := net.ListenTCP("tcp", tcpAddr)

	if err1 != nil {
		log.Fatal("Listen fail")
	} else {
		fmt.Println("listen success")
		fmt.Println("reg address at server:", tcpAddr.IP, tcpAddr.Port)
	}
	for {
		conn, err2 := ln.Accept()
		n := conn.RemoteAddr()
		fmt.Println(n.Network())
		fmt.Println(n.String())
		if err2 != nil {
			log.Fatal("Accept fail")
		} else {
			fmt.Println("accept success")
		}
		go routine(conn)
	}
}

func routine(conn net.Conn) {
	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message))
		// sample process for string received
		newmessage := strings.ToUpper(message)
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
