package greedy

import "math"

func maxSubArray(nums []int) int {
	maxV := nums[0]
	cumV := nums[0]
	for i := 1; i < len(nums); i++ {
		if cumV < 0 {
			cumV = nums[i]
		} else {
			cumV += nums[i]
		}
		if cumV > maxV {
			maxV = cumV
		}
	}
	return maxV
}

func maxSubArrayV2(nums []int) int {
	sum := 0
	maxSum := math.MinInt

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > maxSum {
			maxSum = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return maxSum
}
