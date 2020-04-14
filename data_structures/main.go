package main

import (
	"fmt"
	"struct/structures"
)

func main() {
	var s = structures.NewStack(5)
	s.Push(3)
	s.Push(5)
	fmt.Println("Stack example: ", s)

	var q = structures.NewQueue(7)
	q.Enqueue(1)
	q.Enqueue(3)
	q.Enqueue(5)
	fmt.Println("Queue example: ", q)

	var l = structures.NewList()
	var n_1 = structures.NewNode(7)
	var n_2 = structures.NewNode(9)
	l.List_insert(n_1)
	l.List_insert(n_2)
	fmt.Println("Linked List example : ", l.List_search(7)) // return pointer

}
