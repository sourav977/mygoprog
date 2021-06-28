package main

import "fmt"

//Constants are declared like variables, but with the const keyword as a prefix
//Constants can be character, string, boolean, or numeric values.
//Constants cannot be declared using the := syntax.

//Golang doesn’t have a char data type. It uses byte and rune to represent character values

//untyped const
const r1 = 'a'
const r2 = "a"
const pi = 3.145

//typed const, explicitly specify the type in the declaration
const by byte = 45
const bl bool = true
const i int16 = 17 >> 4
const s string = "sourav"

//const random := "zozo" , can not declared by :=

func main() {
	fmt.Printf("%T\t %v\n", r1, r1) //int32    97
	fmt.Printf("%T\t %v\n", r2, r2) //string   a
	fmt.Printf("%T\t %v\n", pi, pi) //float64  3.145
	fmt.Printf("%T\t %v\n", by, by) //uint8    45

	fmt.Printf("%c\n", by)          // - , character representation of the byte.
	fmt.Printf("%T\t %v\n", bl, bl) //bool     true
	fmt.Printf("%T\t %v\n", i, i)   //int16    1
	fmt.Printf("%T\t %v\n", s, s)   //string   sourav
}

// can you add a const field in struct ? give a try
