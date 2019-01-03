package main

import (
	"fmt"
)

var sl4 []int

func main() {
	sl1 := make([]int, 5, 5)
	fmt.Println(sl1)
	sl1 = append(sl1, 11, 22, 33)
	fmt.Println(sl1)

	var sl2 []int
	fmt.Println(sl2)
	sl2 = append(sl2, 100, 200, 300)
	fmt.Println(sl2)
	sl3 := sl2[1:len(sl2)]
	fmt.Println("sl3:", sl3)

	fmt.Println("sl4:", sl4)
	sl4 = append(sl4, 99, 88, 77)
	fmt.Println("after append, sl4:", sl4)
	func1()
	fmt.Println(sl4)
}
func func1() {
	sl4 = append(sl4, 66, 55, 44)
	fmt.Println("in func1, sl4:", sl4)
}
