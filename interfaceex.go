package main

import "fmt"

type Shaper interface {
	Area() float32
}
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

func main() {
	sq1 := new(Square)
	fmt.Println(sq1)
	sq1.side = 5
	var n1 Square
	fmt.Printf("%v\n", n1)
	// var areaIntf Shaper
	// areaIntf = sq1
	// shorter, without separate declaration:
	// areaIntf := Shaper(sq1)
	// or even:
	areaIntf := sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area())
	rec1 := new(Rectangle)
	rec1.length = 23.50
	rec1.breadth = 12.89
	s := []Shaper{rec1, sq1}
	for _, v := range s {
		result := v.Area()
		fmt.Println(result)
	}
}
