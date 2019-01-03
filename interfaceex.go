package main

import "fmt"

type Shaper interface {
	Area() float32
}
type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
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
}
