package linked_list

import "leetcode/src/define/mylinkednode"

func findMiddleElement(head *mylinkednode.ListNode) *mylinkednode.ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
