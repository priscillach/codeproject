package linked_list

import (
	"leetcode/src/define/mylinkednode"
	"math"
)

func deleteDuplicates82(head *mylinkednode.ListNode) *mylinkednode.ListNode {
	newHead := &mylinkednode.ListNode{
		Next: head,
		Val:  math.MinInt,
	}
	cur := newHead
	for cur != nil {
		if next, isDuplicate := deleteDuplicatesCore82(cur.Next); isDuplicate {
			cur.Next = next
			continue
		}
		cur = cur.Next
	}
	return newHead.Next
}

func deleteDuplicatesCore82(head *mylinkednode.ListNode) (*mylinkednode.ListNode, bool) {
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

func deleteDuplicates82V2(head *mylinkednode.ListNode) *mylinkednode.ListNode {
	newHead := &mylinkednode.ListNode{
		Next: head,
	}

	cur := newHead
	for cur != nil {
		next, isDuplicated := findNext(cur.Next)
		if isDuplicated {
			cur.Next = next
			continue
		}
		cur = cur.Next
	}
	return newHead.Next
}

func findNext(head *mylinkednode.ListNode) (*mylinkednode.ListNode, bool) {
	var isDuplicated bool
	if head == nil {
		return nil, isDuplicated
	}
	cur := head.Next
	for cur != nil {
		if cur.Val == head.Val {
			isDuplicated = true
			cur = cur.Next
		} else {
			return cur, isDuplicated
		}
	}

	return cur, isDuplicated
}
