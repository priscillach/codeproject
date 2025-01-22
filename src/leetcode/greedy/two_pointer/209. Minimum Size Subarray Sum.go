package two_pointer

// https://leetcode.com/problems/move-zeroes/description/
func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	sum := 0
	minLen := 0
	for right < len(nums) {
		for sum < target && right < len(nums) {
			sum += nums[right]
			right++
		}
		for left <= right && sum >= target {
			sum -= nums[left]
			left++
		}
		if left != 0 || right != len(nums) {
			if minLen == 0 || right-left+1 < minLen {
				minLen = right - left + 1
			}
		}
	}
	return minLen
}

func minSubArrayLenV2(target int, nums []int) int {
	left, right := 0, 0
	sum := 0
	minLen := 0
	for ; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target {
			if minLen == 0 || right-left+1 < minLen {
				minLen = right - left + 1
			}
			sum -= nums[left]
			left++
		}
	}
	return minLen
}
