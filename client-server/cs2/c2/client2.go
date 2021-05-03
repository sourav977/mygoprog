package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	var t net.TCPAddr
	t.IP = net.ParseIP("127.0.0.1")
	t.Port = 8081
	conn, _ := net.DialTCP("tcp", nil, &t)
	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, text+"\n")
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
		keepalive := true
		err := conn.SetKeepAlive(keepalive)
		fmt.Println(err)
	}
}
