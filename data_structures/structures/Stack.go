package structures

import "fmt"

type Stack struct {
	//n   int // you have len(array)
	S   []int
	top int
}

func NewStack(n int) Stack {
	stack := new(Stack)
	stack.S = make([]int, n)
	stack.top = -1 //instead 0
	return *stack
}

func Stack_empty(stack Stack) bool {
	return stack.top == -1
}

func (stack *Stack) Push(x int) {
	if stack.top == len(stack.S)-1 {
		fmt.Println("Error overflow")
	} else {
		stack.top += 1
		stack.S[stack.top] = x
	}
}

func (stack *Stack) Pop() int {
	if Stack_empty(*stack) {
		fmt.Println("Error underflow")
		return 0
	} else {
		stack.top -= 1
	}
	return stack.S[stack.top+1]
}

/*
func main() {

	var c = newStack(5)

	fmt.Println(stack_empty(c))
	fmt.Println(c.pop())
	fmt.Println(c) // { array, s.top (array starts in 0) }
	c.push(1)
	c.push(2)
	c.push(3)
	c.push(4)
	c.push(5)
	fmt.Println(c)
	c.push(10)

}
*/
