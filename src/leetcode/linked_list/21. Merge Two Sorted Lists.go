package linked_list

import (
	"leetcode/src/define/mylinkednode"
)

func mergeTwoLists(list1 *mylinkednode.ListNode, list2 *mylinkednode.ListNode) *mylinkednode.ListNode {
	cur1 := list1
	cur2 := list2
	head := &mylinkednode.ListNode{}
	cur3 := head
	for cur1 != nil && cur2 != nil {
		var prev *mylinkednode.ListNode
		if cur1.Val > cur2.Val {
			prev = cur2
			cur2 = cur2.Next
		} else {
			prev = cur1
			cur1 = cur1.Next
		}
		prev.Next = nil
		cur3.Next = prev
		cur3 = cur3.Next
	}
	if cur1 != nil {
		cur3.Next = cur1
	}
	if cur2 != nil {
		cur3.Next = cur2
	}
	return head.Next
}

func mergeTwoListsV2(list1 *mylinkednode.ListNode, list2 *mylinkednode.ListNode) *mylinkednode.ListNode {
	cur1 := list1
	cur2 := list2
	head := &mylinkednode.ListNode{}
	cur3 := head
	for cur1 != nil && cur2 != nil {
		if cur1.Val > cur2.Val {
			cur3.Next = cur2
			cur2 = cur2.Next
		} else {
			cur3.Next = cur1
			cur1 = cur1.Next
		}
		cur3 = cur3.Next
	}
	if cur1 != nil {
		cur3.Next = cur1
	}
	if cur2 != nil {
		cur3.Next = cur2
	}
	return head.Next
}
