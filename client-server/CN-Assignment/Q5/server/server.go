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
	"fmt"
	"io/ioutil"
	"net"
)

func main() {
	// Listen for incoming connections.
	fmt.Println("Launching server...")
	l, err := net.Listen("tcp", "127.0.0.1:8082")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue
		}
		server := conn.LocalAddr()
		client := conn.RemoteAddr()
		//checking
		fmt.Println("Server's network type:", server.Network())
		fmt.Println("Server's IPAddress:Port :", server.String())
		fmt.Println("Client's network type:", client.Network())
		fmt.Println("Client's IPAddress:Port :", client.String())
		// Handle the connection in a new goroutine.
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	// Read the incoming file from the client. Read all file content at once.
	fileBytes, err := ioutil.ReadAll(conn)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}
	// Display the contents of the file.
	fmt.Println("File contents:\n", string(fileBytes))
	// Close the connection when done.
	conn.Close()
}
