package two_pointer

// https://leetcode.com/problems/move-zeroes/description/
func moveZeroes(nums []int) {
	zero, nonZero := 0, 0
	for zero < len(nums) && nonZero < len(nums) {
		if nums[zero] == 0 && nums[nonZero] != 0 && zero < nonZero {
			nums[zero], nums[nonZero] = nums[nonZero], nums[zero]
		}
		for zero < len(nums) && nums[zero] != 0 {
			zero++
		}
		for nonZero < len(nums) && (nums[nonZero] == 0 || nonZero <= zero) {
			nonZero++
		}
	}
}
