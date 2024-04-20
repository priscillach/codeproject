package mybinarayindexedtree

import (
	"fmt"
	"leetcode/src/utils"
)

type BIT struct {
	tree []int
	arr  []int
}

func NewBIT(size int) *BIT {
	// tree[x] means the sum from arr[x - lowbit(i) + 1 : x]
	return &BIT{tree: make([]int, size+1)}
}

func NewBITWithArray(arr []int) *BIT {
	b := NewBIT(len(arr))
	for i := 1; i <= len(arr); i++ {
		b.Add(i, arr[i-1])
	}
	b.arr = arr
	return b
}

// note all index starts from 1

func (b *BIT) Add(index, delta int) {
	// the father range sum of tree[x] is tree[x + lowbit(x)]
	if len(b.arr) > 0 {
		if index > len(b.arr) {
			fmt.Println("arr length invalid")
		}
		b.arr[index-1] += delta
	}
	for i := index; i < len(b.tree); i += utils.LowBit(i) {
		b.tree[i] += delta
	}
}

func (b *BIT) Update(index, val int) {
	if len(b.arr)+1 != len(b.tree) {
		fmt.Println("please use the BIT by NewInitialArray(arr)")
		return
	}
	b.Add(index, val-b.arr[index-1])
	b.arr[index-1] = val
}

// GetPrefixSum the prefix sum of arr[1 : x], note the index start from 1, and x included
func (b *BIT) GetPrefixSum(index int) int {
	sum := 0
	// the next composition range [x - lowbit(i) + 1 : x] of tree[x] is tree[x - lowbit(x)] for sum of arr[1 : x]
	for i := index; i > 0; i -= utils.LowBit(i) {
		sum += b.tree[i]
	}
	return sum
}

func (b *BIT) GetSumByRange(leftIndex int, rightIndex int) int {
	if leftIndex > rightIndex || rightIndex > len(b.tree) || leftIndex < 1 {
		return 0
	}
	return b.GetPrefixSum(rightIndex) - b.GetPrefixSum(leftIndex-1)
}
