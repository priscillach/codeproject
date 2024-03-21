package myheap

import (
	"container/heap"
	"leetcode/src/define/mylinkednode"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MinLinkedNodeHeap []*mylinkednode.ListNode

func (h MinLinkedNodeHeap) Len() int           { return len(h) }
func (h MinLinkedNodeHeap) Less(i, j int) bool { return h[i].Val <= h[j].Val }
func (h MinLinkedNodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinLinkedNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*mylinkednode.ListNode))
}

func (h *MinLinkedNodeHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func HeapSortDesc(arr []int) {
	h := &MaxHeap{}
	heap.Init(h)

	for _, x := range arr {
		heap.Push(h, x)
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = heap.Pop(h).(int)
	}
}

func HeapSortAsc(arr []int) {
	h := &MinHeap{}
	heap.Init(h)

	for _, x := range arr {
		heap.Push(h, x)
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = heap.Pop(h).(int)
	}
}
