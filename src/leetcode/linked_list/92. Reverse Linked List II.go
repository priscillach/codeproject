package linked_list

import (
	"leetcode/src/define/mylinkednode"
	"math"
)

func ReverseBetween(head *mylinkednode.ListNode, left int, right int) *mylinkednode.ListNode {
	if left == right {
		return head
	}
	var cur *mylinkednode.ListNode

	// sentinel node
	newHead := &mylinkednode.ListNode{
		Val:  math.MaxInt64,
		Next: head,
	}
	prev := head
	var last *mylinkednode.ListNode
	for i := 0; i < right; i++ {
		// the first node of reverse list
		if i == left-1 {
			cur = prev
		}
		// the last node of non-reverse list 1
		if i == left-2 {
			last = prev
		}

		// if reverse all list, the new head is the last node
		if left-1 == 0 && i == right-1 {
			newHead.Next = prev
		}
		prev = prev.Next
	}
	// prev now is the first node of non-reverse list 2
	for i := 0; i < right-left+1; i++ {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	// prev now is the first node of reverse list

	// if last is not nil, last.Next should be the first node of reverse list
	if last != nil {
		last.Next = prev
	}
	return newHead.Next
}

func reverseBetweenV2(head *mylinkednode.ListNode, left int, right int) *mylinkednode.ListNode {
	newHead := &mylinkednode.ListNode{Next: head}

	cur := head
	cnt := 0

	leftTail := newHead
	var rightFirst *mylinkednode.ListNode

	for {
		if cnt == left-2 {
			leftTail = cur
		}
		if cnt == right {
			rightFirst = cur
			break
		}
		cnt++
		cur = cur.Next
	}
	prev := rightFirst
	cur = leftTail.Next
	for i := 0; i < right-left+1; i++ {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	leftTail.Next = prev
	return newHead.Next
}
