package linked_list

import "leetcode/src/define/mylinkednode"

// https://leetcode.com/problems/linked-list-cycle-ii/description/
func detectCycle(head *mylinkednode.ListNode) *mylinkednode.ListNode {
	var hasCycleFlag bool
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			hasCycleFlag = true
			break
		}
	}
	if !hasCycleFlag {
		return nil
	}
	for head != slow {
		head = head.Next
		slow = slow.Next
	}
	return head
}
