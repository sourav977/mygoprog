package main

import (
	"fmt"
)

func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from:", r)
	}
}

func fullName(firstName, lastName *string) {
	defer recoverName()
	if firstName == nil || lastName == nil {
		panic("runtime error: first name or last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func main() {
	defer fmt.Println("deferred call in main")
	fname := "sourav"
	//lname := "patnaik"
	//fullName(&fname, &lname)
	fullName(&fname, nil)
	fmt.Println("returned normally from main")
}
