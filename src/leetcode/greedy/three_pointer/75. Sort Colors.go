package three_pointer

func sortColors(nums []int) {
	left, mid, right := 0, 0, len(nums)-1
	for mid <= right {
		if nums[mid] == 0 {
			nums[left], nums[mid] = nums[mid], nums[left]
			left++
			mid++
		} else if nums[mid] == 1 {
			mid++
		} else {
			nums[mid], nums[right] = nums[right], nums[mid]
			right--
		}
	}
}
