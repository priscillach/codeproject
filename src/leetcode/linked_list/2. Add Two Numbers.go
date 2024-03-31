package linked_list

import "leetcode/src/define/mylinkednode"

func addTwoNumbers(l1 *mylinkednode.ListNode, l2 *mylinkednode.ListNode) *mylinkednode.ListNode {
	newHead := &mylinkednode.ListNode{}
	cur := newHead
	carry := 0
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if sum+carry > 9 {
			sum = (sum + carry) % 10
			carry = 1
		} else {
			sum = sum + carry
			carry = 0
		}
		cur.Next = &mylinkednode.ListNode{Val: sum}
		cur = cur.Next
	}
	if carry > 0 {
		cur.Next = &mylinkednode.ListNode{Val: carry}
	}
	return newHead.Next
}
