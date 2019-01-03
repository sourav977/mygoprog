package main

import "fmt"

type Foo struct {
	A int
	B string
	C interface{}
}

type Bar struct {
	A []int
}

const (
	a1 = iota
	a2 = iota
	a3 = iota
	a4
)

func main() {
	a := map[string]bool{"A": true, "B": false}
	var b map[string]bool
	//copy(b, a) //.\copymap.go:8:6: arguments to copy must be slices; have map[string]bool, map[string]bool
	fmt.Println(a, b)

	c := Foo{A: 1, B: "one", C: "two"}
	d := Foo{A: 1, B: "one", C: "two"}
	fmt.Println(c == d)

	//e := Bar{A: []int{1}}
	//f := Bar{A: []int{1}}
	//fmt.Println(e == f) .\copymap.go:27:16: invalid operation: e == f (struct containing []int cannot be compared)

	var g interface{}
	var h interface{}

	g = 10
	h = 10
	fmt.Println(g == h)

	// x := []int{1, 2, 3}
	// y := []int{1, 2, 3}
	// fmt.Println(x == y)
	fmt.Println(a1, a2, a3, a4)
}
