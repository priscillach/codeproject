package linked_list

import "leetcode/src/define/mylinkednode"

func sortList(head *mylinkednode.ListNode) *mylinkednode.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	tmp := slow.Next
	slow.Next = nil
	head1 := sortList(head)
	head2 := sortList(tmp)
	return mergeTwoListsV2(head1, head2)
}
