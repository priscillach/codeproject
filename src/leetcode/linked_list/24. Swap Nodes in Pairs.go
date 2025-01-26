package linked_list

import "leetcode/src/define/mylinkednode"

// https://leetcode.com/problems/swap-nodes-in-pairs/description/
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

func swapPairsV2(head *mylinkednode.ListNode) *mylinkednode.ListNode {
	newHead := &mylinkednode.ListNode{
		Next: head,
	}
	cur := newHead
	for cur.Next != nil && cur.Next.Next != nil {
		tail := cur.Next.Next.Next
		first := cur.Next
		cur.Next = cur.Next.Next
		cur.Next.Next = first
		first.Next = tail
		cur = cur.Next.Next
	}
	return newHead.Next
}
