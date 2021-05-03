package main

import (
	"fmt"
)

var i int

func main() {
	func1()
	fmt.Println(i)
}

func func1() {
	defer i = 30 // .\deferFunc2.go:13:10: expression in defer must be function call, compiletime error
}