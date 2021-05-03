package main

import (
	"fmt"
	"time"
)

func main() {
	//var ch chan int
	//fmt.Println(len(ch), cap(ch))
	//fmt.Print("%v\n", ch)
	//ch <- 10
	//fmt.Println(len(ch), cap(ch))

	ch1 := make(chan int)
	fmt.Println(len(ch1), cap(ch1))
	ch1 <- 10
	time.Sleep(1)
	fmt.Println(<-ch1)
	fmt.Println(len(ch1), cap(ch1))
}
