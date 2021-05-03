/*
first run server.go in a new terminal.
then run client.go in another terminal.

you can run client.go more then one terminal, server will
accept all client request running on different terminals.

press ctrl + c  to exit
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, err := net.Listen("tcp", ":8081")
	fmt.Println("ln:", ln)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	// run loop forever (or until ctrl-c)
	for {
		conn, err := ln.Accept()
		fmt.Println("conn:", conn)
		if err != nil {
			log.Fatal(err)
		}
		x := conn.LocalAddr()
		y := conn.RemoteAddr()
		//checking
		fmt.Println("local network type:", x.Network())
		fmt.Println("local ip:port :", x.String())
		fmt.Println("remote network type:", y.Network())
		fmt.Println("remote ip:port :", y.String())
		go func(c net.Conn) {
			// will listen for message to process ending in newline (\n)
			message, _ := bufio.NewReader(conn).ReadString('\n')
			// output message received
			fmt.Print("Message Received:", string(message))
			// sample process for string received
			newmessage := strings.ToUpper(message)
			// send new string back to client
			c.Write([]byte(newmessage + "\n"))
			defer c.Close()
		}(conn)
	}
}
