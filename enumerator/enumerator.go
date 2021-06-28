package main

import (
	"fmt"
)

//An enum, or enumerator, is a data type consisting of a set of named constant values
//In Go, enumerated constants are created using the iota enumerator.
//The iota keyword represents successive integer constants 0, 1, 2 ..
//It resets to 0 whenever the word const appears in the source code
//constants created at compile time

//iota is
//const iota untyped int = 0
const (
	C0 = iota
	C1 = iota
	C2 = iota
)
const (
	C3 = iota
	C4
	C5
)

const (
	C6 = iota
	_  //blank identifier to skip a value
	C8
	C9
)

const (
	C10 = iota + 100
	C11
	C12
)

const (
	C13 = iota - 100
	C14
	C15
)

const (
	C16 = 0
	C17
	C18
)

const (
	C19 = iota / 100
	C20
	C21
)

func main() {
	fmt.Println(C0, C1, C2)    // 0 1 2
	fmt.Println(C3, C4, C5)    // 0 1 2
	fmt.Println(C6, C8, C9)    // 0 2 3
	fmt.Println(C10, C11, C12) // 100 101 102
	fmt.Println(C13, C14, C15) // ?
	fmt.Println(C16, C17, C18) // ? and why
	fmt.Println(C19, C20, C21) // ? and why
}
