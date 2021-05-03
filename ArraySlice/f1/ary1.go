package main

import (
	"fmt"
)

type str struct {
	name string
	age  int
}

func main() {
	a := new(str)
	fmt.Println(a)
	fmt.Println(a.name, a.age)
	var b str
	fmt.Println(b)
	fmt.Println(b.name, b.age)

	sl1 := make([]int, 5)
	fmt.Println(sl1)
	fmt.Println(len(sl1), cap(sl1))
	sl1 = append(sl1, 11, 22, 33, 44, 55)
	fmt.Println(sl1)
	mp1 := make(map[string]int)
	fmt.Println(mp1)
}
