package linked_list

import "leetcode/src/define/mylinkednode"

func getIntersectionNode(headA, headB *mylinkednode.ListNode) *mylinkednode.ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a := headA
	b := headB

	for a != b {
		if a == nil {
			a = headB
			continue
		}
		if b == nil {
			b = headA
			continue
		}
		a = a.Next
		b = b.Next
	}
	return a
}
