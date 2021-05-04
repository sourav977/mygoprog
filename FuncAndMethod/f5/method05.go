package main

import (
	"fmt"
	"strings"
	"time"
)

type mystr string

func (s mystr) UpperCase() string {
	return strings.ToUpper(string(s))
}
func main() {
	var s mystr = "sourav"
	go fmt.Println(s.UpperCase())
	time.Sleep(1)

}
