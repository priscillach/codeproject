package linked_list

import "leetcode/src/define/mylinkednode"

// input: 1->8->3->6->5->4->7->2->NULL
// output: 1->2->3->4->5->6->7->8->NULL

func sortOddAscEvenDescList(root *mylinkednode.ListNode) *mylinkednode.ListNode {
	if root == nil {
		return nil
	}
	head1 := root
	head2 := root.Next

	odd := head1
	even := head2
	for even != nil {
		nextOdd := even.Next
		var nextEven *mylinkednode.ListNode
		if even.Next != nil {
			nextEven = even.Next.Next
		}
		odd.Next = nextOdd
		even.Next = nextEven
		odd = nextOdd
		even = nextEven
	}
	reverseHead2 := ReverseList(head2)
	return mergeTwoListsV2(head1, reverseHead2)
}
