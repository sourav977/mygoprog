package main

import (
	"fmt"
)

func main() {
	func1(1, 2, 3, 4, 5)

}

func func1(nums ...int) {
	fmt.Println(nums[2])
}
