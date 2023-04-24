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
	"net"
	"os"
)

func main() {
	fmt.Println("Launching client...")

	// connect to Server running at 127.0.0.1:8081
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	client := conn.LocalAddr()
	server := conn.RemoteAddr()
	//checking
	fmt.Println("Client's network type:", client.Network())
	fmt.Println("Client's IPAddress:Port :", client.String())
	fmt.Println("Server's network type:", server.Network())
	fmt.Println("Server's IPAddress:Port :", server.String())

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Text to send: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	// send to socket
	fmt.Fprintf(conn, text+"\n")
	fmt.Print("Closing Client")

}
