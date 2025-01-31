package linked_list

import "leetcode/src/define/mylinkednode"

func isPalindrome(head *mylinkednode.ListNode) bool {
	var prev, slow, fast *mylinkednode.ListNode = nil, head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}
	// if the length of the linked list is odd, then make slow is the (len / 2 + 1) th node
	if fast != nil && fast.Next == nil {
		slow = slow.Next
	}
	for prev != nil {
		if prev.Val != slow.Val {
			return false
		}
		prev = prev.Next
		slow = slow.Next
	}
	return true
}

func isPalindromeV2(head *mylinkednode.ListNode) bool {
	var fast, slow, prev *mylinkednode.ListNode = head, head, nil
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	// odd num
	if fast != nil {
		slow = slow.Next
	}

	for prev != nil && slow != nil && prev.Val == slow.Val {
		prev = prev.Next
		slow = slow.Next
	}

	return prev == nil && slow == nil
}
