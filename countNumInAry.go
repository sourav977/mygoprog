package main

import (
	"fmt"
)

// count the occurance of each element in an array
func main() {
	ary := [10]int{1, 2, 1, 2, 3, 1, 2, 2, 3, 4}
	fmt.Println(ary)
	mymap := make(map[int]int)

	for i := 0; i < 10; i++ {
		mymap[ary[i]]++
	}
	fmt.Println(mymap)
}
