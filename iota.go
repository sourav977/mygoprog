package main

import (
	"fmt"
)

const (
	a, b, c = iota, iota, iota
	d       = iota
	g
)

const (
	e = iota
	f
)

func main() {
	fmt.Println(a, b, c, d, g, e, f)
}
