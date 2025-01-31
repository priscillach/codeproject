package two_pointer

func removeDuplicates(nums []int) int {
	left := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		left++
		nums[left] = nums[i]
	}
	return left + 1
}
