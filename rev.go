package main

import "fmt"

func main() {
	s := "sourav"
	fmt.Println(len(s))
	var reverse string
	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	fmt.Println(reverse)
}
