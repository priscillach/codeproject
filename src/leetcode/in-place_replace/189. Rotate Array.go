package in_place_replace

import "leetcode/src/utils/arrayhelper"

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n                                   // 防止 k > n
	arrayhelper.ReverseBetween(nums, 0, n-1) // 1. 反转整个数组
	arrayhelper.ReverseBetween(nums, 0, k-1) // 2. 反转前 k 个元素
	arrayhelper.ReverseBetween(nums, k, n-1) // 3. 反转剩余元素
}

func rotateV2(nums []int, k int) {
	n := len(nums)
	k %= n     // 防止 k > n
	count := 0 // 记录已经移动的元素个数

	for start := 0; count < n; start++ {
		current := start
		prev := nums[start]
		for {
			next := (current + k) % n
			nums[next], prev = prev, nums[next] // 交换值
			current = next
			count++
			if start == current { // 一轮循环后回到起点
				break
			}
		}
	}
}
