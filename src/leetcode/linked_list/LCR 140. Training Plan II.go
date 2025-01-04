package linked_list

import (
	"leetcode/src/define/mylinkednode"
)

// https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/description/
// finish times: 1
func trainingPlan(head *mylinkednode.ListNode, cnt int) *mylinkednode.ListNode {
	first := head
	for i := 0; i < cnt-1; i++ {
		first = first.Next
	}
	second := head
	for first != nil && first.Next != nil {
		first = first.Next
		second = second.Next
	}
	return second
}
