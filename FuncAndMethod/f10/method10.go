package main

import "fmt"

type User struct {
	Name   string
	Period int
}

//method with pointer reciver
func (u *User) Work() {
	fmt.Println(u.Name, "has worked for", u.Period, "hrs.")
}

//method with value reciver
// func (u User) Work() {
// 	fmt.Println(u.Name, "has worked for", u.Period, "hrs.")
// }

func main() {
	uval := User{"UserVal", 5}
	uval.Work()

	pval := &User{"UserPtr", 6}
	pval.Work()
}
