package linked_list

import "leetcode/src/define/mylinkednode"

func addTwoNumbers445(l1 *mylinkednode.ListNode, l2 *mylinkednode.ListNode) *mylinkednode.ListNode {
	var stack1 []*mylinkednode.ListNode
	var stack2 []*mylinkednode.ListNode
	cur := l1
	for cur != nil {
		stack1 = append(stack1, cur)
		cur = cur.Next
	}
	cur = l2
	for cur != nil {
		stack2 = append(stack2, cur)
		cur = cur.Next
	}

	var carry int
	var head *mylinkednode.ListNode
	for len(stack1) > 0 || len(stack2) > 0 {
		sum := carry
		newHead := &mylinkednode.ListNode{Next: head}
		if len(stack1) > 0 {
			sum += stack1[len(stack1)-1].Val
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) > 0 {
			sum += stack2[len(stack2)-1].Val
			stack2 = stack2[:len(stack2)-1]
		}
		newHead.Val = sum % 10
		carry = sum / 10
		head = newHead

	}
	if carry > 0 {
		head = &mylinkednode.ListNode{Val: carry, Next: head}
	}
	return head
}
