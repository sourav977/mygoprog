//Data serialisation
package main

import (
	"encoding/asn1"
	"fmt"
	"net"
	"os"
	"time"
)

//This client and server are exchanging ASN.1 encoded data values, not textual strings.
// go run asn1client.go
func main() {
	service := "127.0.0.1:1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	fmt.Println(tcpAddr)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		daytime := time.Now()
		// Ignore return network errors.
		mdata, _ := asn1.Marshal(daytime)
		conn.Write(mdata)
		conn.Close() // we're finished
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
