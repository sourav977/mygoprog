package main

import (
	"fmt"
)

func main() {
	ary := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("initial ary:", ary)
	sl1 := make([]int, 6, 7)
	sl1 = ary[2:8]
	fmt.Println("initial sl1:", sl1)
	for j := 0; j < len(sl1); j++ {
		sl1[j] += 11
	}
	fmt.Println("after first modification sl1:", sl1)
	fmt.Println("after first modification ary:", ary)
	func1(ary)
	fmt.Println("after func1 call, ary:", ary)
	for i := 0; i < len(ary); i++ {
		ary[i] += 4
	}
	fmt.Println("new ary:", ary)
	fmt.Println("new slice:", sl1)

	sl2 := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}
	sl3 := sl2[2:8]
	fmt.Println("initial sl2:", sl2)
	fmt.Println("initial sl3:", sl3)
	for j := 0; j < len(sl3); j++ {
		sl3[j] += 100
	}
	fmt.Println("after sl2:", sl2)
	fmt.Println("after sl3:", sl3)
	func2(sl2)
	fmt.Println("after func2 call, sl2:", sl2)
	fmt.Println("after func2 call, sl3:", sl3)
	sl2 = append(sl2, 599, 699, 799, 899)
	fmt.Println("after append sl2:", sl2)
	fmt.Println("after append sl3:", sl3)

}
func func1(a [10]int) {
	fmt.Println("in func1 before ary:", a)
	for i := 0; i < len(a); i++ {
		a[i] += 200
	}
	fmt.Println("in func1 after ary:", a)
}

func func2(a []int) {
	fmt.Println("in func2 before slc2:", a)
	for i := 0; i < len(a); i++ {
		a[i] += 300
	}
	fmt.Println("in func2 after slc2:", a)
	a = append(a, 499, 509, 519, 529)
	fmt.Println("a:", a)
}
