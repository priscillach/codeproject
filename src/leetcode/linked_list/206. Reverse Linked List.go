package linked_list

import "leetcode/src/define/linkednode"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *linkednode.ListNode) *linkednode.ListNode {
	_, newHead := reverse(head)
	return newHead

}

func reverse(head *linkednode.ListNode) (last, newHead *linkednode.ListNode) {
	if head == nil || head.Next == nil {
		return head, head
	}
	last, newHead = reverse(head.Next)
	head.Next = nil
	last.Next = head
	return head, newHead

}
