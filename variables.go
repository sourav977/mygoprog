package main

import "fmt"

func main() {
	var n1 uint8 // Unsigned 8-bit integers (0 to 255)
	//n1 = 255 + 1 //constant 256 overflows uint8
	n1 = 255
	fmt.Println(n1)

	var n2 int8
	n2 = 130 //  constant 130 overflows int8
	fmt.Println(n2)

	var n3 float32
	n3 = -17.89
	fmt.Println(n3)
}
