package main

import (
	"container/list"
	"fmt"
)

func main() {
	node1 := list.New()
	node1.PushBack(10)
	node1.PushBack(4)
	node1.PushBack(12)
	node1.PushBack(8)

	node2 := list.New()
	node2.PushBack(5)
	node2.PushBack(10)
	node2.PushBack(13)
	node2.PushFront(1)
	node2.PushBack(15)

	fmt.Println("node1: ", node1)
	fmt.Println("node2: ", node2)

	for i := node2.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	sortLinkedList(node1)
	sortLinkedList(node2)
	fmt.Println("after sort, node1: ")
	for i := node1.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	fmt.Println("after sort, node2: ")
	for i := node2.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}

func sortLinkedList(node *list.List) *list.List {
	current := node.Front()
	if node.Front() == nil {
		return nil
	} else {
		for current != nil {
			index := current.Next()
			for index != nil {
				if current.Value.(int) > index.Value.(int) {
					temp := current.Value
					current.Value = index.Value
					index.Value = temp
				}
				index = index.Next()
			}
			current = current.Next()
		}
	}

	return node
}
