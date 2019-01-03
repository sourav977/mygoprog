package main

import (
	"fmt"
)

func main() {
	s := "Sourav"
	// fmt.Println(s)
	// s[2] = "z"
	fmt.Println(s)
	b := []byte(s)
	fmt.Println(len(b))
	b[0] = 'G'
	s = string(b)
	fmt.Println(s)
	fmt.Print("%t\n%v\n", 'A', 'A')
	s1 := "hi"
	s1 = "bye"
	fmt.Println(s1)

	//swap 3 variables
	var x, y, z int = 10, 20, 30
	fmt.Println(x, y, z)
	x, y, z = y, z, x
	fmt.Println(x, y, z)

	ch := make(chan string)
	fmt.Println(len(ch), cap(ch))

}
