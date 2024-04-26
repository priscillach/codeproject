package two_pointer

func moveZeroes(nums []int) {
	zero, nonZero := 0, 0
	for zero < len(nums) && nonZero < len(nums) {
		if nums[zero] == 0 && nums[nonZero] != 0 && zero < nonZero {
			nums[zero], nums[nonZero] = nums[nonZero], nums[zero]
		}
		if zero < len(nums) && nums[zero] != 0 {
			zero++
		}
		if nonZero < len(nums) && (nums[nonZero] == 0 || nonZero <= zero) {
			nonZero++
		}
	}
}
