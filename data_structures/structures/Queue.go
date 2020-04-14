package structures

import "fmt"

type Queue struct {
	head int
	tail int
	Q    []int
	n    int // number of elements used
}

func NewQueue(n int) Queue {
	q := new(Queue)
	q.Q = make([]int, n)
	return *q
}

func Empty_queue(queue Queue) bool {
	return queue.n == 0
}

func Full_queue(queue Queue) bool {
	return queue.n == len(queue.Q)
}

func (queue *Queue) Enqueue(x int) {
	if Full_queue(*queue) {
		fmt.Println("Error full queue")
	} else {
		queue.Q[queue.tail] = x
		if queue.tail == len(queue.Q)-1 {
			queue.tail = 0
		} else {
			queue.tail += 1
		}
		queue.n += 1
	}
}

func (queue *Queue) Dequeue() int {
	if Empty_queue(*queue) {
		fmt.Println("Error empty queue")
		return 0
	} else {
		x := queue.Q[queue.head]
		queue.n -= 1
		if queue.head == len(queue.Q)-1 {
			queue.head = 0
		} else {
			queue.head += 1
		}
		return x
	}

}

/*
func main() {
	a := newQueue(5)
	fmt.Println(empty_queue(a))
	fmt.Println(a.dequeue())
	a.enqueue(2)
	a.enqueue(4)
	fmt.Println(a)
	fmt.Println(a.dequeue())
	a.enqueue(3)
	a.enqueue(5)
	a.enqueue(7)
	a.enqueue(10)
	fmt.Println(a)
	a.enqueue(100)
}*/
