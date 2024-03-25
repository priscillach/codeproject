package linked_list

import "leetcode/src/define/mylinkednode"

func ReverseList(head *mylinkednode.ListNode) *mylinkednode.ListNode {
	var prev, cur *mylinkednode.ListNode = nil, head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func ReverseListV2(head *mylinkednode.ListNode) *mylinkednode.ListNode {
	_, newHead := reverse(head)
	return newHead
}

func reverse(head *mylinkednode.ListNode) (last, newHead *mylinkednode.ListNode) {
	if head == nil || head.Next == nil {
		return head, head
	}
	last, newHead = reverse(head.Next)
	head.Next = nil
	last.Next = head
	return head, newHead

}
