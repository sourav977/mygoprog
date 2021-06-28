package main

import "fmt"

/*
Global variables are defined outside of a function, usually on top of the program.
Global variables hold their value throughout the lifetime of the program and they
can be accessed inside any of the functions defined for the program
*/

var global0 bool
var global1 = 32.4
var global2 int
var global3 string = "sourav"

func main() {
	fmt.Printf("%T\t %v\n", global0, global0)     //bool     false
	fmt.Printf("%T\t %v\n", global1, global1)     //float64  32.4
	fmt.Printf("%T\t %v\n", global2, global2+100) //int      100
	fmt.Printf("%T\t %v\n", global3, global3)     //string   sourav

	global2 += 500
	fmt.Printf("%T\t %v\n", global2, global2) //int      500
}
