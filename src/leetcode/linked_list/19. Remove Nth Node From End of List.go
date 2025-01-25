package linked_list

import "leetcode/src/define/mylinkednode"

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
func removeNthFromEnd(head *mylinkednode.ListNode, n int) *mylinkednode.ListNode {
	newHead := &mylinkednode.ListNode{Next: head}
	first := newHead
	for i := 0; i < n; i++ {
		first = first.Next
	}
	second := newHead
	for first != nil && first.Next != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return newHead.Next
}

func removeNthFromEndV2(head *mylinkednode.ListNode, n int) *mylinkednode.ListNode {
	newHead := &mylinkednode.ListNode{Next: head}
	fast, slow := newHead, newHead
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return newHead.Next
}
