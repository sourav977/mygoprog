package main

import (
	"fmt"
)

func main() {
	var sl1 []int
	fmt.Println("len:", len(sl1), "cap:", cap(sl1))
	sl1 = []int{1, 2}
	fmt.Println("len:", len(sl1), "cap:", cap(sl1))
	fmt.Println(sl1)
	func1(&sl1)
	fmt.Println("in main func, after func1 call, sl1:", sl1)

	sl2 := make([]int, 2, 4)
	sl2 = []int{11, 22}
	func2(&sl2)
	fmt.Println("in main func, after func2 call, sl2:", sl2)
	var ar1 [5][3]int
	fmt.Println("len:", len(ar1), "cap:", cap(ar1))
	var ar2 [7]int
	fmt.Println("len:", len(ar2), "cap:", cap(ar2))

}

func func1(a *[]int) {
	*a = append(*a, 3, 4, 5)
	fmt.Println("in func1, after append, sl1:", a)
}

func func2(b *[]int) {
	b1 := *b
	b1[0] = 100
	*b = append(*b, 44, 55, 66)
	fmt.Println("in func2, after append, sl2:", b)
}
