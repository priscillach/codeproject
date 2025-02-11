package other

import (
	"container/heap"
	"sort"
)

type MedianFinder struct {
	minHeap MinHeap
	maxHeap MaxHeap
}

type MinHeap struct {
	sort.IntSlice
}
type MaxHeap struct {
	sort.IntSlice
}

func (h *MaxHeap) Less(i, j int) bool {
	return h.IntSlice.Less(j, i)
}

func (h *MaxHeap) Push(x any) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *MaxHeap) Pop() any {
	x := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return x
}

func (h *MinHeap) Push(x any) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *MinHeap) Pop() any {
	x := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return x
}

func Constructor295() MedianFinder {
	finder := MedianFinder{}
	heap.Init(&finder.minHeap)
	heap.Init(&finder.maxHeap)
	return finder
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(&this.maxHeap, num)
	heap.Push(&this.minHeap, heap.Pop(&this.maxHeap))
	if this.minHeap.Len() > this.maxHeap.Len() {
		heap.Push(&this.maxHeap, heap.Pop(&this.minHeap))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxHeap.Len() > this.minHeap.Len() {
		return float64(this.maxHeap.IntSlice[0])
	}
	return float64(this.maxHeap.IntSlice[0]+this.minHeap.IntSlice[0]) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
