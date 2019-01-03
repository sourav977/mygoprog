package main

import (
	"fmt"
	"sort"
)

func main() {

	a := make([]int, 5)

	fmt.Scan(&a[0], &a[1], &a[2], &a[3], &a[4])
	sort.Ints(a)
	min := a[0] + a[1] + a[2] + a[3]
	max := a[1] + a[2] + a[3] + a[4]
	fmt.Println(min, max)
}
