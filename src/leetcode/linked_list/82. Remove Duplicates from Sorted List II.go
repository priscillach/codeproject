package linked_list

import (
	"leetcode/src/define/mylinkednode"
	"math"
)

func deleteDuplicates(head *mylinkednode.ListNode) *mylinkednode.ListNode {
	newHead := &mylinkednode.ListNode{
		Next: head,
		Val:  math.MinInt,
	}
	cur := newHead
	for cur != nil {
		if next, isDuplicate := detectDuplicates(cur.Next); isDuplicate {
			cur.Next = next
			continue
		}
		cur = cur.Next
	}
	return newHead.Next
}

func detectDuplicates(head *mylinkednode.ListNode) (*mylinkednode.ListNode, bool) {
	if head == nil {
		return nil, false
	}
	val := head.Val
	cnt := 0
	head = head.Next
	for head != nil {
		if head.Val != val {
			break
		}
		head = head.Next
		cnt++
	}
	return head, cnt > 0
}