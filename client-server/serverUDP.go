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

	var t net.UDPAddr
	t.IP = net.ParseIP("127.0.0.1")
	t.Port = 8081
	// listen on all interfaces
	ln, err1 := net.ListenUDP("udp", &t)
	if err1 != nil {
		log.Fatal("Listen fail")
	} else {
		fmt.Println("listen success")
	}
	for {
		go routine(ln)
	}
	// var b []byte
	// conn.Read(b)
	// fmt.Println(string(b))
	// conn.Write([]byte("got your msg"))
	//conn.Close()

}

func routine(conn *net.UDPConn) {
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
