package linked_list

import (
	"leetcode/src/define/linkednode"
	"math"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *linkednode.ListNode, left int, right int) *linkednode.ListNode {
	if left == right {
		return head
	}
	var cur *linkednode.ListNode
	newHead := &linkednode.ListNode{
		Val:  math.MaxInt64,
		Next: head,
	}
	prev := head
	var last *linkednode.ListNode
	for i := 0; i < right; i++ {
		if i == left-1 {
			cur = prev
		}
		if i == left-2 {
			last = prev
		}
		if left-1 == 0 && i == right-1 {
			newHead.Next = prev
		}
		prev = prev.Next
	}

	for i := 0; i < right-left+1; i++ {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	if last != nil {
		last.Next = prev
	}
	return newHead.Next
}
