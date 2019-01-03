package main

import (
	"fmt"
)

func main() {
	map1 := make(map[string]int)
	map2 := make(map[int]string)
	map1 = map[string]int{"a": 10, "b": 20}
	map2 = map[int]string{100: "apple", 200: "orange"}

	sl1 := make([]map[string]int, 3)
	sl1 = []map[string]int{map1}
	fmt.Println(sl1)
	fmt.Println(map2)
}
