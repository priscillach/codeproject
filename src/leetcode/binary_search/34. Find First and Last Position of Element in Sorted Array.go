package binarysearch

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	// first get the left target
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if nums[left] != target {
		return []int{-1, -1}
	}
	tmp := left
	// first get the right target
	left, right = 0, len(nums)-1
	for left < right {
		mid := left + (right-left+1)>>1
		if nums[mid] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return []int{tmp, left}
}
