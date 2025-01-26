package linked_list

import "leetcode/src/define/mylinkednode"

func deleteDuplicates83(head *mylinkednode.ListNode) *mylinkednode.ListNode {
	newHead := &mylinkednode.ListNode{}
	newHead.Next = head
	for head != nil {
		head.Next = deleteDuplicatesCore83(head)
		head = head.Next
	}
	return newHead.Next
}

func deleteDuplicatesCore83(head *mylinkednode.ListNode) *mylinkednode.ListNode {
	if head == nil {
		return nil
	}
	val := head.Val
	cur := head.Next

	for cur != nil {
		if cur.Val != val {
			break
		}
		cur = cur.Next
	}
	return cur
}

func deleteDuplicatesV2(head *mylinkednode.ListNode) *mylinkednode.ListNode {
	if head == nil {
		return nil
	}
	cur := head
	var next *mylinkednode.ListNode
	for cur != nil {
		next = cur.Next
		for next != nil && next.Val == cur.Val {
			next = next.Next
		}
		cur.Next = next
		cur = next
	}
	return head
}
