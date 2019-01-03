package main

import "fmt"

func func1() int {
	defer fmt.Println("hello")
	//fmt.Println(z)
	return 10
}

func main() {
	x := func1()
	fmt.Println(x)

}
