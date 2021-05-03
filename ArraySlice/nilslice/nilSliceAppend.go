package main

import (
	"fmt"
)

func main() {
	var sl1 []int
	fmt.Println(len(sl1), cap(sl1), sl1)
	sl1 = append(sl1, 1, 2, 3, 4, 5)
	fmt.Println(sl1)
	fmt.Println(len(sl1), cap(sl1), sl1)

	sl2 := make([]int, 0)
	fmt.Println(len(sl2), cap(sl2), sl2)
	sl2 = append(sl2, 100, 200)
	fmt.Println(len(sl2), cap(sl2), sl2)

	sl3 := make([]int, 3)
	fmt.Println(len(sl3), cap(sl3), sl3)
	sl3 = append(sl3, 100, 200)
	fmt.Println(len(sl3), cap(sl3), sl3)
}
