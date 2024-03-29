package linked_list

import (
	"container/heap"
	"leetcode/src/define/myheap"
	"leetcode/src/define/mylinkednode"
)

func mergeKLists(lists []*mylinkednode.ListNode) *mylinkednode.ListNode {
	newHead := &mylinkednode.ListNode{}
	cur := newHead
	h := &myheap.MinLinkedNodeHeap{}
	heap.Init(h)
	for _, head := range lists {
		if head != nil {
			heap.Push(h, head)
		}
	}
	for h.Len() > 0 {
		cur.Next = heap.Pop(h).(*mylinkednode.ListNode)
		cur = cur.Next
		if cur.Next != nil {
			heap.Push(h, cur.Next)
		}
	}
	return newHead.Next
}

func mergeKListsV2(lists []*mylinkednode.ListNode) *mylinkednode.ListNode {
	if len(lists) == 0 {
		return nil
	}
	newHead := lists[0]
	for _, list := range lists[1:] {
		newHead = mergeTwoListsV2(newHead, list)
	}
	return newHead
}
