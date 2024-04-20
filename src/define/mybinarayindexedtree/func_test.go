package mybinarayindexedtree

import (
	"fmt"
	"testing"
)

func TestBit(t *testing.T) {
	arr := []int{1, 3, 2, -5, 6, 4, 2, 1}

	// Create a new Fenwick Tree
	ft := NewBITWithArray(arr)
	// Build the tree by updating each element

	// Example queries
	fmt.Println("Prefix sum from index 1 to 4:", ft.GetPrefixSum(4))
	fmt.Println("Prefix sum from index 3 to 7:", ft.GetPrefixSum(7)-ft.GetPrefixSum(2))
	fmt.Println("Prefix sum from index 3 to 7:", ft.GetSumByRange(3, 7))

	// Add an element at index 3, arr[3] += 2
	ft.Add(3, 2)
	// Example query after update
	fmt.Println("Prefix sum from index 1 to 4 after add:", ft.GetPrefixSum(4))

	ft.Update(3, 1)
	// Example query after add
	fmt.Println("Prefix sum from index 1 to 4 after update:", ft.GetPrefixSum(4))
}
