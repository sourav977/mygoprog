package main

import "fmt"

type stack []int

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(data int) {
	*s = append((*s), data)
}

func (s *stack) pop() {
	if s.isEmpty() {
		fmt.Println("Stack is empty.")
	} else {
		index := len(*s) - 1
		pop := (*s)[index]
		fmt.Printf("%d popped\n", pop)
		*s = (*s)[:index]
	}
}

func main() {
	var s stack
	//push 10 elements
	for i := 10; i > 0; i-- {
		s.push(i)
	}
	fmt.Println("After Push, Elements in Stack: ", s)
	fmt.Println("Stack length ", len(s), " capacity ", cap(s))
	//pop 10 elements
	for len(s) > 0 {
		s.pop()
	}
	fmt.Println("After Pop, Elements in Stack: ", s)
	fmt.Println("Stack length ", len(s), " capacity ", cap(s))
	s.pop() //possible?
}
