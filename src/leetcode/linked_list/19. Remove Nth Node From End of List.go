package linked_list

import "leetcode/src/define/mylinkednode"

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
