package linked_list

import "leetcode/src/define/linkednode"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *linkednode.ListNode, k int) *linkednode.ListNode {
	node, cnt := head, 0
	for cnt < k {
		if node == nil {
			return head
		}
		node = node.Next
		cnt++
	}
	prev := reverseKGroup(node, k)
	cur := head
	for cnt > 0 {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
		cnt--
	}
	return prev
}
