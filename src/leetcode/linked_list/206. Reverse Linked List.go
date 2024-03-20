package linked_list

import "leetcode/src/define/mylinkednode"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *mylinkednode.ListNode) *mylinkednode.ListNode {
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
