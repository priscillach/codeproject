package linked_list

import "leetcode/src/define/mylinkednode"

// https://leetcode.com/problems/reorder-list/description/
func reorderList(head *mylinkednode.ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	cur := slow.Next
	slow.Next = nil
	cur1 := head
	cur2 := ReverseList(cur)
	for cur1 != nil && cur2 != nil {
		cur1, cur2 = insertNodeAfter(cur1, cur2)
	}
}

func insertNodeAfter(left, node *mylinkednode.ListNode) (*mylinkednode.ListNode, *mylinkednode.ListNode) {
	if left == nil || node == nil {
		return nil, nil
	}
	next1 := left.Next
	left.Next = node
	next2 := node.Next
	node.Next = next1
	return next1, next2
}

func reorderListV2(head *mylinkednode.ListNode) {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	cur := slow.Next
	slow.Next = nil
	var prev *mylinkednode.ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	secondHead := prev
	cur = head
	for cur != nil && secondHead != nil {
		next := cur.Next
		cur.Next = secondHead
		secondNext := secondHead.Next
		secondHead.Next = next
		secondHead = secondNext
		cur = next
	}
}
