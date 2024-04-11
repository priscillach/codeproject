package list

import "testing"

func TestPermute(t *testing.T) {
	nums := []int{5, 4, 6, 2}
	permute(nums)
}

func TestPermuteUnique(t *testing.T) {
	nums := []int{1, 1, 2}
	permuteUnique(nums)
}

func TestCombinationSum(t *testing.T) {
	combinationSum([]int{2, 3, 6, 7}, 7)
}
