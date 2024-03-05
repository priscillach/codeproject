package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	cur1 := list1
	cur2 := list2
	head := &ListNode{}
	cur3 := head
	for cur1 != nil && cur2 != nil {
		var prev *ListNode
		if cur1.Val > cur2.Val {
			prev = cur2
			cur2 = cur2.Next
		} else {
			prev = cur1
			cur1 = cur1.Next
		}
		prev.Next = nil
		cur3.Next = prev
		cur3 = cur3.Next
	}
	if cur1 != nil {
		cur3.Next = cur1
	}
	if cur2 != nil {
		cur3.Next = cur2
	}
	return head.Next
}

func mergeTwoListsV2(list1 *ListNode, list2 *ListNode) *ListNode {
	cur1 := list1
	cur2 := list2
	head := &ListNode{}
	cur3 := head
	for cur1 != nil && cur2 != nil {
		if cur1.Val > cur2.Val {
			cur3.Next = cur2
			cur2 = cur2.Next
		} else {
			cur3.Next = cur1
			cur1 = cur1.Next
		}
		cur3 = cur3.Next
	}
	if cur1 != nil {
		cur3.Next = cur1
	}
	if cur2 != nil {
		cur3.Next = cur2
	}
	return head.Next
}
