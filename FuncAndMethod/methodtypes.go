package main

import (
	"fmt"
	"time"
)

// .\methodtypes.go:7:6: func init must have no arguments and no return values
func init() {
	fmt.Println("hello from init")
	go func1() //Goroutines may be run during program initialization
	time.Sleep(1)
}

func func1() {
	fmt.Println("go routine called from init() ")
}

type intary [2]int        // creating alias
type intslc []int         // creating alias
type mymap map[string]int // creating alias

func (i intary) addition() { // receiver type is an array
	res := 0
	for x := 0; x < len(i); x++ {
		res += i[x]
	}
	fmt.Println(res)
}

func (s intslc) addition() { // receiver type is a slice
	res := 0
	for x := 0; x < len(s); x++ {
		res += s[x]
	}
	fmt.Println(res)
}

func (m mymap) addition() { // receiver type is a map
	res := 0
	for _, v := range m {
		res += v
	}
	fmt.Println(res)
}

func main() {
	fmt.Println("hello from main")
	var inary intary
	inary[0] = 10
	inary[1] = 20
	inary.addition()

	var inslc intslc
	inslc = append(inslc, 100, 200, 300, 400, 500)
	inslc.addition()

	var mp mymap
	fmt.Printf("%t, %V\n", mp, mp)
	mp = make(mymap, 3)
	fmt.Printf("%T, %v\n", mp, mp)
	mp["a"] = 300
	mp["b"] = 400
	mp["c"] = 32
	fmt.Printf("%T, %v\n", mp, mp)
	mp.addition()
	//x := 10
	//y := 12.89
	//z := x + y // .\methodtypes.go:69:9: invalid operation: x + y (mismatched types int and float64)
	//fmt.Println(z)
}
