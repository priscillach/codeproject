package mylinkednode

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := new(ListNode)
	cur := head
	for _, num := range nums {
		cur.Next = &ListNode{
			Val: num,
		}
		cur = cur.Next
	}
	return head.Next
}
