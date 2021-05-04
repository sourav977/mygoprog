package main

import "fmt"

type Square struct {
	side float32
}
type Rectangle struct {
	length, breadth float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (rec *Rectangle) Area() float32 {
	return rec.length * rec.breadth
}

type Shaper interface {
	Area() float32
}

type Container interface {
	Shaper
}

// func (c Container) Caller() {
// 	fmt.Println(c.Area())
// }
func main() {
	var n1 Square
	n1.side = 30.2
	var s1 Shaper = &n1
	fmt.Println(s1.Area())

	var n2 = &Rectangle{20, 30}
	var s2 Shaper = n2
	fmt.Println(s2.Area())

}
