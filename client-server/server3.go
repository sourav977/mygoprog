package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// only needed below for sample processing

func main() {

	fmt.Println("Launching server...")
	// listen on all interfaces
	ln, err1 := net.Listen("tcp", "127.0.0.1:8001")
	if err1 != nil {
		log.Fatal("Listen fail")
	} else {
		fmt.Println("listen success")
	}
	// accept connection on port
	// conn, err2 := ln.Accept()
	// if err2 != nil {
	// 	log.Fatal("Accept fail")
	// } else {
	// 	fmt.Println("accept success")
	// }
	for {
		conn, err2 := ln.Accept()
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
