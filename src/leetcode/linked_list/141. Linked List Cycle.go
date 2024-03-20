package linked_list

import (
	"leetcode/src/define/mylinkednode"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *mylinkednode.ListNode) bool {
	one := head
	two := head
	for two != nil && two.Next != nil {
		one = one.Next
		two = two.Next.Next
		if one == two {
			return true
		}
	}
	return false
}
