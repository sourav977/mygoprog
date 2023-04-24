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
	fmt.Println("Launching client...")
	// Connect to the server.
	conn, err := net.Dial("tcp", "127.0.0.1:8082")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()
	client := conn.LocalAddr()
	server := conn.RemoteAddr()
	//checking
	fmt.Println("Client's network type:", client.Network())
	fmt.Println("Client's IPAddress:Port :", client.String())
	fmt.Println("Server's network type:", server.Network())
	fmt.Println("Server's IPAddress:Port :", server.String())

	// Read the file to be sent.
	fileBytes, err := ioutil.ReadFile("2022mt12211.txt")
	if err != nil {
		fmt.Println("Error reading file:", err.Error())
		return
	}

	// Send the file to the server.
	_, err = conn.Write(fileBytes)
	if err != nil {
		fmt.Println("Error sending file:", err.Error())
		return
	}
}
