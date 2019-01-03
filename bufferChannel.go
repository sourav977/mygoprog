package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2) //buffer channel
	ch <- 1
	ch <- 2
	//ch <- 3 //fatal error: all goroutines are asleep - deadlock!
	fmt.Println("len:", len(ch), "cap:", cap(ch))
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println("len:", len(ch), "cap:", cap(ch))
	//fmt.Println(<-ch) // fatal error: all goroutines are asleep - deadlock!
}
