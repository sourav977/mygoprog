package main

import (
	"github.com/reiver/go-telnet"
)

func main() {

	var handler telnet.Handler = telnet.EchoHandler

	err := telnet.ListenAndServe(":5555", handler)
	if nil != err {
		//@TODO: Handle this error better.
		panic(err)
	}
}
