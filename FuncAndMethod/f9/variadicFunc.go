package main

import (
	"fmt"
	"unsafe"
)

func main() {
	print("hello")
	println("sourav patnaik")
	var n = new([3]int)
	println(len(n), cap(n), "%v\n", n)
	println(n[0])
	println(unsafe.Sizeof(n))

	var action int
	fmt.Println("Enter 1 for Student and 2 for Professional")
	fmt.Scanln(&action)
	defer func() {
		action := recover()
		fmt.Println(action)
	}()
	/*  Use of Switch Case in Golang */
	switch action {
	case 1:
		fmt.Printf("I am a  Student")
	case 2:
		fmt.Printf("I am a  Professional")
	default:
		panic(fmt.Sprintf("I am a  %d", action))
	}
}
