package main

import (
	"fmt"
)

// fscan used for file scan
func main() {
	var ary [5]int
	fmt.Println("enter 5 space separated num:")
	for i := 0; i < 5; i++ {
		// 	//fmt.Fscanf(os.Stdin, "%d", &ary[i])
		// 	//fmt.Fscanln(os.Stdin, &ary[i])
		// 	//fmt.Fscan(os.Stdin, &ary[i])
		fmt.Scan(&ary[i])
	}
	//fmt.Scanln(&ary)
	//fmt.Fscanln(os.Stdin, &ary)
	fmt.Println(ary)

	x, y := "sourav", 24
	fmt.Println(x, y)
	z1 := 23.78
	z2 := 4
	z = z1 + z2
	fmt.Println(z)
}
