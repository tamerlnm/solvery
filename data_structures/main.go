package main

import (
	"datastructures/ds"
	"fmt"
)

func main() {
	// Stack
	stack := ds.Stack{}
	stack.Push(1)
	stack.Push(2)

	fmt.Println("Stack:", stack)

	top, ok := stack.Peek()
	if ok {
		fmt.Println("Stack Top:", top)
	} else {
		fmt.Println("Stack is empty")
	}

	popped, ok := stack.Pop()
	if ok {
		fmt.Println("Stack Pop:", popped)
	} else {
		fmt.Println("Stack is empty")
	}

	fmt.Printf("Stack: %v\n\n", stack)

	// Queue
	queue := ds.Queue{}
	queue.Enqueue(10)
	queue.Enqueue(20)

	fmt.Println("Queue:", queue)

	front, ok := queue.Dequeue()
	if ok {
		fmt.Println("Queue Front:", front)
	} else {
		fmt.Println("Queue is empty")
	}
	fmt.Printf("Queue: %v\n", queue)
}
