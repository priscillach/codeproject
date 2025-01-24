package algo

import "math/rand"

// FastSort fast sort
func FastSort(nums []int) {
	FastSortByRange(nums, 0, len(nums)-1)
}

func FastSortByRange(nums []int, left, right int) {
	if left >= right {
		return
	}
	cur := FastSortCore(nums, left, right)
	FastSortByRange(nums, left, cur-1)
	FastSortByRange(nums, cur+1, right)
}

func FastSortCore(nums []int, left, right int) int {
	randIdx := left + rand.Intn(right-left+1)
	tmp := nums[randIdx]
	nums[randIdx] = nums[right]
	for left < right {
		for left < right && nums[left] <= tmp {
			left++
		}
		nums[right] = nums[left]
		for left < right && nums[right] >= tmp {
			right--
		}
		nums[left] = nums[right]
	}
	nums[right] = tmp
	return left
}
