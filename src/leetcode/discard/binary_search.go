package leetcode

// 有重复的取最后一个
func biSearchV1(nums []int, k int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low+1)>>1
		if nums[mid] > k {
			high = mid - 1
		} else {
			low = mid
		}
	}
	if nums[low] == k {
		return low
	}
	return -1
}

// 有重复的取第一个
func biSearchV2(nums []int, k int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)>>1
		if nums[mid] < k {
			low = mid + 1
		} else {
			high = mid
		}
	}
	if nums[low] == k {
		return low
	}
	return -1
}
