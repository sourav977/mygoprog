package main

import (
	"fmt"
)

func main() {
	var sl1 []int
	sl1 = append(sl1, 1, 2, 3, 4, 5)
	fmt.Println(sl1)

	sl2 := make([]int, 2)
	sl2 = append(sl2, 100, 200)
	fmt.Println(len(sl2), cap(sl2), sl2)
}
