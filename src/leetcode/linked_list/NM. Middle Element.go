package linked_list

import "leetcode/src/define/mylinkednode"

// the right when length is even
func findRightMiddleElement(head *mylinkednode.ListNode) *mylinkednode.ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// the left when length is even
func findLeftMiddleElement(head *mylinkednode.ListNode) *mylinkednode.ListNode {
	slow, fast := head, head
	if fast != nil {
		fast = fast.Next
	}
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
