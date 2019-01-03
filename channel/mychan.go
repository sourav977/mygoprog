package main

import "fmt"
import "time"

const h = "hello world"

func main() {
	var c1 chan int
	fmt.Printf("%T %v\n", c1, c1) //chan int <nil>
	c2 := make(chan int)
	fmt.Printf("%T %v\n", c2, c2) //chan int 0xc04204a060
	fmt.Printf("%T\n", h)
	done := make(chan bool)
	go hello(done)
	time.Sleep(3)
	data := <-done
	fmt.Println("ok main", data)
}

func hello(done chan bool) {
	fmt.Println("before write")
	done <- true
	time.Sleep(1)
	fmt.Println("after")
}
