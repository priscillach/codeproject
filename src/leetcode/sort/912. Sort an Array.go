package sort

import "golang.org/x/exp/rand"

func sortArray(nums []int) []int {
	sort(nums, 0, len(nums)-1)
	return nums
}

func sort(nums []int, left, right int) {
	if left >= right {
		return
	}
	cur := fastSort(nums, left, right)
	sort(nums, left, cur-1)
	sort(nums, cur+1, right)
}

func fastSort(nums []int, left, right int) int {
	randIdx := left + rand.Intn(right-left+1)
	tmp := nums[randIdx]
	nums[randIdx] = nums[right]
	nums[right] = tmp
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