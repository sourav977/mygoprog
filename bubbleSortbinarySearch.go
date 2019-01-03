package main

import (
	"fmt"
)

//bubble sort
func main() {
	var ary [6]int
	fmt.Println("enter 6 space separated num:")
	for i := 0; i < 6; i++ {
		fmt.Scan(&ary[i])
	}
	fmt.Println("input unsorted ary:", ary)
	for i := 0; i < 6; i++ {
		for j := 0; j < 6-i-1; j++ {
			if ary[j] > ary[j+1] {
				ary[j+1], ary[j] = ary[j], ary[j+1]
			}
		}
	}
	fmt.Println("(bubble sort) sorted ary:", ary)

	//binary search
	flag := 0
	fmt.Println("enter the no to search:")
	var s int
	fmt.Scan(&s)
	low := 0
	high := len(ary)
	index := 0
	for low <= high {
		mid := (low + high) / 2
		if s == ary[mid] {
			flag = 1
			index = mid
			break
		} else {
			if s > ary[mid] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	if flag == 1 {
		fmt.Println("(binary search)element found, at position:", index+1)
	} else {
		fmt.Println("element not found")
	}
}
