//package general
//go run: cannot run non-main package
package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := struct {
		Foo string
		Bar int
	}{"foo", 2}

	v := reflect.ValueOf(x)

	values := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	fmt.Println(values)
}

// func main(){
// 	//32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems
// 	var b1 bool
// 	var b2, b3 bool
// 	var i1, i2, i3 int

// }
