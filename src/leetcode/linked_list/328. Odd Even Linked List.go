package linked_list

import "leetcode/src/define/mylinkednode"

func oddEvenList(head *mylinkednode.ListNode) *mylinkednode.ListNode {
	newHead1 := &mylinkednode.ListNode{Next: head}
	newHead2 := &mylinkednode.ListNode{}
	odd := head
	even := newHead2
	prev := newHead1
	for odd != nil && odd.Next != nil {
		next := odd.Next
		odd.Next = odd.Next.Next
		prev = odd
		odd = odd.Next
		even.Next = next
		next.Next = nil
		even = even.Next
	}
	if odd == nil {
		prev.Next = newHead2.Next
		return newHead1.Next
	} else {
		odd.Next = newHead2.Next
	}
	return newHead1.Next
}
