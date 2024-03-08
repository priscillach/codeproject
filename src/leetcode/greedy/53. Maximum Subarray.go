package greedy

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
