package main

import "fmt"

type User struct {
	Name   string
	Period int
}

type Worker interface {
	Work()
}

/*
./method11.go:20:8: cannot use uval (type User) as type Worker in argument to DoWork:
        User does not implement Worker (Work method has pointer receiver)
*/
func (u *User) Work() {
	fmt.Println(u.Name, "has worked for", u.Period, "hrs.")
	u.Period = 20
}

// func (u User) Work() {
// 	fmt.Println(u.Name, "has worked for", u.Period, "hrs.")
// }

func main() {
	// uval := User{"UserVal", 5}
	// DoWork(uval)

	pval := &User{"UserPtr", 6}
	DoWork(pval)
	fmt.Println(pval)
}

func DoWork(w Worker) {
	w.Work()
}
