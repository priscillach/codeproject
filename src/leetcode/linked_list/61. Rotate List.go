package linked_list

import (
	"leetcode/src/define/mylinkednode"
)

// https://leetcode.com/problems/rotate-list/description/
func rotateRight(head *mylinkednode.ListNode, k int) *mylinkednode.ListNode {
	fast := head
	cnt := 0
	for fast != nil {
		fast = fast.Next
		cnt++
	}
	if cnt == 0 {
		return head
	}
	fast = head

	var prev *mylinkednode.ListNode
	for i := 0; i <= k%cnt; i++ {
		prev = fast
		fast = fast.Next
	}
	slow := head

	for fast != nil {
		prev = fast
		fast = fast.Next
		slow = slow.Next
	}

	next := slow.Next
	if next == nil {
		return head
	}
	if prev == nil {
		return head
	}
	slow.Next = nil
	prev.Next = head
	return next
}
