package main

import "fmt"

func main() {
	p1 := [2]int{1, 2}
	p2 := [2]int{1, 2}
	p3 := &[2]int{3, 4}
	p4 := p3
	//p3 := &[2]int{1, 2}
	if p1 == p2 {
		fmt.Println("p1 == p2")
	} else {
		fmt.Println("p1 != p2")
	}
	if p3 == p4 {
		fmt.Println("p3 == p4")
	} else {
		fmt.Println("p3 != p4")
	}
	i := a()
	fmt.Println("i:", i)
	//invalid operation: p2 == p3 (mismatched types [2]int and *[2]int)
	// if p2 == p3 {
	// 	fmt.Println("p2 == p3")
	// } else {
	// 	fmt.Println("p2 != p3")
	// }
}

func a() int {
	i := 0
	defer fmt.Println(i)
	i++
	return i
}
