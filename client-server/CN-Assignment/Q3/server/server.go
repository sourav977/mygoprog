/*
first run server.go in a new terminal.
then run client.go in another terminal.

you can run client.go more then one terminal, server will
accept all client request running on different terminals.

press ctrl + c  to exit
*/

/*
2022MT12211
Sourav Patnaik
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, err := net.Listen("tcp", ":8081") //start Server on 127.0.0.1:8081
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	// run loop forever (or until ctrl-c)
	for {
		conn, err := ln.Accept()
		fmt.Println("Connection Accepted:", conn.LocalAddr().String())
		if err != nil {
			log.Fatal(err)
		}
		server := conn.LocalAddr()
		client := conn.RemoteAddr()
		//checking
		fmt.Println("Server's network type:", server.Network())
		fmt.Println("Server's IPAddress:Port :", server.String())
		fmt.Println("Client's network type:", client.Network())
		fmt.Println("Client's IPAddress:Port :", client.String())
		go func(c net.Conn) {
			// will listen for message to process ending in newline (\n)
			message, _ := bufio.NewReader(conn).ReadString('\n')
			// print received message from client at server side
			fmt.Print("Message Received from Client:", string(message))
			defer c.Close()
		}(conn)
	}
}
