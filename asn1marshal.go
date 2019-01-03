package main

import (
	"encoding/asn1"
	"fmt"
	"os"
)

//dealing with ASN.1 encoded data values, not textual strings.
func main() {
	mdata, err := asn1.Marshal("sourav")
	checkError(err)
	var n string
	_, err1 := asn1.Unmarshal(mdata, &n)
	checkError(err1)
	fmt.Println("After marshal/unmarshal: ", n)

}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
