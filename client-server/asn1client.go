//Data serialisation
package main

import (
	"encoding/asn1"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"time"
)

//This client and server are exchanging ASN.1 encoded data values, not textual strings.
// go run asn1client.go localhost:1200
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	conn, err := net.Dial("tcp", service)
	checkError(err)
	defer conn.Close()
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	var newtime time.Time
	_, err1 := asn1.Unmarshal(result, &newtime)
	checkError(err1)
	fmt.Println("After marshal/unmarshal: ", newtime.String())
	os.Exit(0)
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
