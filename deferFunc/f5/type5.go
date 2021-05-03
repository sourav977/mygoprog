package main

import "fmt"

func main() {
	var x = 10
	fmt.Println("hi")
	defer func() {
		y := func1(&x)
		fmt.Println(y)
	}()
	fmt.Println("done")
	//op
	//hi
	//done
	//20
}

func func1(x *int) int {
	*x += 10
	return *x
}
