package main

import (
	"fmt"
)

func main() {
	sl1 := []int{10, 20, 30}
	fmt.Println("slice before call:", sl1)
	modify1(sl1) //call by value slice
	fmt.Println("slice after call:", sl1)
	//modify2(&sl1)
	apnd(&sl1)
	fmt.Println("after apnd:", sl1)

	ar1 := [3]int{101, 201, 301}
	ar2 := [3]int{101, 201}
	if ar1 == ar2 {
		fmt.Println("ok ary")
	} else {
		fmt.Println("not ok ary")
	}
	fmt.Println("array before call:", ar1)
	modify2(ar1)
	fmt.Println("array after call:", ar1)
	sl2 := []int{8, 20, 30}
	if sl1[0] == sl2[0] {
		fmt.Println("ok slice")
	} else {
		fmt.Println("not ok slice")
	}
	if ar1 == ar2 {
		fmt.Println("ok ary")
	} else {
		fmt.Println("not ok ary")
	}

	p1 := struct {
		a string
		b int
	}{"left", 4}
	p2 := struct {
		a string
		b int
	}{a: "righ", b: 11}
	if p1 == p2 {
		fmt.Println("Same struct")
	} else {
		fmt.Println("not same struct")
	}

	pair := [2]int{4, 2}
	ptr := &[2]int{4, 2}
	ptr2 := &pair

	if &pair != ptr {
		fmt.Println("pointing different")
	}
	if &pair == ptr2 {
		fmt.Println("pointing the same")
	}

}
func modify1(n []int) {
	for i := range n {
		n[i] -= 2
	}
	fmt.Println("n:", n)
	n = append(n, 100)
	fmt.Println("n:", n)
}
func apnd(a *[]int) {
	*a = append(*a, 40, 50)
	fmt.Println(a)
}
func modify2(a [3]int) {
	for i := range a {
		a[i] += 7
	}
	fmt.Println("a:", a)
}

// func modify2(m *[]int) {
// 	// for j := range m {
// 	// 	m[j] += 7
// 	// }
// 	// cannot range over m (type *[]int)
// 	&m[0] += 7
// 	&m[1] += 7
// 	&m[2] += 7
// }
