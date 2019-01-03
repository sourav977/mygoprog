package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("bufferChannel.go")
	if err != nil {
		fmt.Println("error in reading file")
		return
	}
	fmt.Println(string(data))
}
