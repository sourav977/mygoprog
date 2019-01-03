package main

import (
	"fmt"
	"sort"
)

func main() {
	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}
	sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("By name:", people)

	sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	fmt.Println("By age:", people)

	sl2 := []int{99, 55, 66, 77, 11, 22, 33, 44, 88}
	//sort.SliceStable(sl2, func(i, j int) bool { return sl2[i] < sl2[j] })
	sort.Sort(sl2)
	fmt.Println("sorted sl2:", sl2)
}
