package linked_list

import "leetcode/src/define/mylinkednode"

func swapPairs(head *mylinkednode.ListNode) *mylinkednode.ListNode {
	newHead := &mylinkednode.ListNode{
		Next: head,
	}
	cur := newHead
	for cur.Next != nil && cur.Next.Next != nil {
		next := cur.Next.Next.Next
		node1 := cur.Next
		node2 := cur.Next.Next
		cur.Next = node2
		node2.Next = node1
		node1.Next = next
		cur = node1
	}
	return newHead.Next
}
