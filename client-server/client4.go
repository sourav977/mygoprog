package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	conn, err1 := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err1)
	if err1 == nil {
		fmt.Println("client connected to server successfully")
		fmt.Println(`type "exit" to dis-connect`)
	}
	keepalive := true
	conn.SetKeepAlive(keepalive)
	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// send to socket

		text += "\n"
		conn.Write([]byte(text))
		//fmt.Fprintf(conn, text+"\n")
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)

	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
