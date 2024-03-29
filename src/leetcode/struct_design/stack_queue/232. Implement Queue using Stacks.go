package stack_queue

type MyQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.stack1 = append(this.stack1, x)
}

func (this *MyQueue) Pop() int {
	if len(this.stack2) == 0 {
		for len(this.stack1) > 0 {
			this.stack2 = append(this.stack2, pop(&this.stack1))
		}
	}
	x := pop(&this.stack2)
	return x
}

func (this *MyQueue) Peek() int {
	if len(this.stack2) == 0 {
		for len(this.stack1) > 0 {
			x := pop(&this.stack1)
			this.stack2 = append(this.stack2, x)
		}
	}
	return this.stack2[len(this.stack2)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.stack1)+len(this.stack2) == 0
}

func pop(stack *[]int) int {
	x := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return x
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
