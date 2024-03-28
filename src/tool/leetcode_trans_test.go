package tool

import (
	"fmt"
	"strings"
	"testing"
)

func TestTransLCP(t *testing.T) {
	transCode := `
func removeNthFromEnd(head *mylinkednode.ListNode, n int) *mylinkednode.ListNode {
	newHead := &mylinkednode.ListNode{Next: head}
	first := newHead
	for i := 0; i < n; i++ {
		first = first.Next
	}
	second := newHead
	for first != nil && first.Next != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return newHead.Next
}
`
	transCode = strings.ReplaceAll(transCode, "mytreenode.", "")
	transCode = strings.ReplaceAll(transCode, "mylinkednode.", "")
	if strings.Contains(transCode, "utils.Max") {
		transCode = strings.ReplaceAll(transCode, "utils.Max", "max")
		transCode += `
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
`
	}
	if strings.Contains(transCode, "utils.Min") {
		transCode = strings.ReplaceAll(transCode, "utils.Min", "max")
		transCode += `
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
`
	}
	fmt.Println(transCode)
}
