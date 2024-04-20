package binary_indexed_tree

import "leetcode/src/utils"

func lengthOfLIS(nums []int) int {
	MAX_SIZE := 20001
	BASE := 10001
	bit := NewBIT(MAX_SIZE)
	for i := 0; i < len(nums); i++ {
		subLongestLength := bit.Query(BASE + nums[i] - 1)
		bit.Update(nums[i]+BASE, subLongestLength+1)
	}
	return bit.Query(MAX_SIZE)
}

// the difference from mybinaryindexedtree.BIT is ?
type BIT struct {
	tree []int
}

func NewBIT(size int) *BIT {
	return &BIT{tree: make([]int, size+1)}
}

func (b *BIT) Update(index, delta int) {
	for i := index; i < len(b.tree); i += utils.LowBit(i) {
		b.tree[i] = utils.Max(delta, b.tree[i])
	}
}

func (b *BIT) Query(index int) int {
	sum := 0
	for i := index; i > 0; i -= utils.LowBit(i) {
		sum = utils.Max(sum, b.tree[i])
	}
	return sum
}
