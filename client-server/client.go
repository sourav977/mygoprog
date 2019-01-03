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
	"net"
	"os"
)

func main() {
	//conn, _ := net.Dial("tcp", "127.0.0.1:8081")  //uncomment to see diff behaviour
	for {
		conn, _ := net.Dial("tcp", "127.0.0.1:8081") //comment to see diff behaviour
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
		// send to socket
		fmt.Fprintf(conn, text+"\n")
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}
