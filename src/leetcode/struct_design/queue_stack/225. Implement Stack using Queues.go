package queue_stack

type MyStackWithTwoQueues struct {
	queue1 []int
	queue2 []int
}

func ConstructorV1() MyStackWithTwoQueues {
	return MyStackWithTwoQueues{}
}

func (this *MyStackWithTwoQueues) Push(x int) {
	this.queue2 = append(this.queue2, x)
	for len(this.queue1) > 0 {
		x := this.queue1[0]
		this.queue1 = this.queue1[1:]
		this.queue2 = append(this.queue2, x)
	}
	this.queue1, this.queue2 = this.queue2, this.queue1
}

func (this *MyStackWithTwoQueues) Pop() int {
	x := this.queue1[0]
	this.queue1 = this.queue1[1:]
	return x
}

func (this *MyStackWithTwoQueues) Top() int {
	x := this.queue1[0]
	return x
}

func (this *MyStackWithTwoQueues) Empty() bool {
	return len(this.queue1) == 0 && len(this.queue2) == 0
}

type MyStackWithOneQueue struct {
	queue []int
}

func ConstructorV2() MyStackWithOneQueue {
	return MyStackWithOneQueue{}
}

func (this *MyStackWithOneQueue) Push(x int) {
	this.queue = append(this.queue, x)
	for i := 0; i < len(this.queue)-1; i++ {
		last := this.queue[0]
		this.queue = this.queue[1:]
		this.queue = append(this.queue, last)
	}
}

func (this *MyStackWithOneQueue) Pop() int {
	x := this.queue[0]
	this.queue = this.queue[1:]
	return x
}

func (this *MyStackWithOneQueue) Top() int {
	x := this.queue[0]
	return x
}

func (this *MyStackWithOneQueue) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
