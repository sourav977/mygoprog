package main

import (
	"fmt"
)

func main() {
	var doClose = false
	//defer func with no argument
	defer func() {
		if doClose {
			fmt.Println("doclose")
		}
	}()
	doClose = true

	//defer func arguments are evaluated at the time when they called
	var i int = 0
	defer func(i int) {
		defer fmt.Println(i)
		return
	}(i)
	i++
}
