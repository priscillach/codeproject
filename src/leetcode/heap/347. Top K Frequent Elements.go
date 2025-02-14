package _heap

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]*Count)
	for _, num := range nums {
		if _, ok := m[num]; !ok {
			m[num] = &Count{
				Num: num,
			}
		}
		m[num].Cnt++
	}

	h := &MyHeap{}
	heap.Init(h)
	for _, v := range m {
		heap.Push(h, v)
	}

	var res []int
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(h).(*Count).Num)
	}
	return res
}

type Count struct {
	Num int
	Cnt int
}

type MyHeap []*Count

func (h MyHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MyHeap) Less(i, j int) bool {
	return h[i].Cnt > h[j].Cnt
}

func (h MyHeap) Len() int {
	return len(h)
}

func (h *MyHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func (h *MyHeap) Push(x any) {
	*h = append(*h, x.(*Count))
}
