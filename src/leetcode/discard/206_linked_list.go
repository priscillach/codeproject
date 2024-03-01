package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	if cur == nil {
		return nil
	}

	for next := head.Next; next != nil; {
		tmp := next.Next
		cur.Next = nil
		next.Next = cur

		cur = next
		next = tmp
	}
	return cur
}
