<<<<<<< Updated upstream
package main

import "fmt"

type queue []int

func (q *queue) isEmpty() bool {
	return len(*q) == 0
}
func (q *queue) enqueue(data int) {
	*q = append((*q), data)
}

func (q *queue) dequeue() {
	if q.isEmpty() {
		fmt.Println("Queue is empty.")
	} else {
		index := len(*q)
		dq := (*q)[0]
		fmt.Printf("%d dequeued\n", dq)
		*q = (*q)[1:index]
	}
}

func main() {
	var qu queue
	//push 10 elements
	for i := 1; i <= 10; i++ {
		qu.enqueue(i)
	}
	fmt.Println("After Enqueue, Elements in Queue: ", qu)
	fmt.Println("Queue length ", len(qu), " capacity ", cap(qu))
	//pop 10 elements
	for len(qu) > 0 {
		qu.dequeue()
	}
	fmt.Println("After Dequeue, Elements in Queue: ", qu)
	fmt.Println("Queue length ", len(qu), " capacity ", cap(qu))
	qu.dequeue() //possible?
}
=======
package main

import "fmt"

type queue []int

func (q *queue) isEmpty() bool {
	return len(*q) == 0
}
func (q *queue) enqueue(data int) {
	*q = append((*q), data)
}

func (q *queue) dequeue() {
	if q.isEmpty() {
		fmt.Println("Queue is empty.")
	} else {
		index := len(*q)
		dq := (*q)[0]
		fmt.Printf("%d dequeued\n", dq)
		*q = (*q)[1:index]
	}
}

func main() {
	var qu queue
	//push 10 elements
	for i := 1; i <= 10; i++ {
		qu.enqueue(i)
	}
	fmt.Println("After Enqueue, Elements in Queue: ", qu)
	fmt.Println("Queue length ", len(qu), " capacity ", cap(qu))
	//pop 10 elements
	for len(qu) > 0 {
		qu.dequeue()
	}
	fmt.Println("After Dequeue, Elements in Queue: ", qu)
	fmt.Println("Queue length ", len(qu), " capacity ", cap(qu))
	qu.dequeue() //possible?
}
>>>>>>> Stashed changes
