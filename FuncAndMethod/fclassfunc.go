package main

import (
	"fmt"
)

func main() {
	a := func(y int) int { //called ananymous func since they don't have name
		fmt.Println("hello world first class function")
		x := 10 / y
		return x
	}
	//a()
	//fmt.Printf("%T", a)
	secondfunc(a)
}

func secondfunc(f func(a int) int) {
	x := f(2)
	fmt.Printf("%T \n %d", f, x)
}
