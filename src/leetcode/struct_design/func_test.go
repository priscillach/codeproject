package struct_design

import (
	"leetcode/src/leetcode/struct_design/stack_queue"
	"testing"
)

func TestMyQueue(t *testing.T) {
	myQueue := stack_queue.Constructor()
	myQueue.Push(1)
	myQueue.Push(2)
	myQueue.Peek()
	myQueue.Pop()
	myQueue.Empty()
}
