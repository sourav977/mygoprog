package main

import "fmt"

func func1() int {
	defer fmt.Println("hello")
	fmt.Println("inside func1")
	return 10
}

func main() {
	x := func1()
	fmt.Println(x)

}
