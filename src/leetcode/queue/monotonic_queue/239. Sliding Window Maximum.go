package monotonic_queue

func maxSlidingWindow(nums []int, k int) []int {
	var deQueue []int
	var res []int
	for i := 0; i < len(nums); i++ {
		for len(deQueue) > 0 && deQueue[0] <= i-k {
			deQueue = deQueue[1:]
		}
		for len(deQueue) > 0 && nums[deQueue[len(deQueue)-1]] <= nums[i] {
			deQueue = deQueue[:len(deQueue)-1]
		}
		deQueue = append(deQueue, i)
		if i >= k-1 {
			res = append(res, nums[deQueue[0]])
		}
	}
	return res
}
