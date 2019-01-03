package main

import (
	"fmt"
	"time"
)

var a int = 1

func hello() int {
	fmt.Println("goroutine")
	a++
	return a
}

type emp struct {
	name string
}

func (e emp) func1(n string) int {
	e.name = n
	fmt.Println("in func1", e)
	return 1
}

func main() {
	go hello()
	time.Sleep(1)
	x := hello()
	fmt.Println(x)
	fmt.Println("main function")
	e1 := new(emp)
	fmt.Println(e1)
	go e1.func1("sourav")
	time.Sleep(1)
	fmt.Println(e1)
}
