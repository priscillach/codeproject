package binarysearch

import "math/rand"

// https://leetcode.com/problems/kth-largest-element-in-an-array/submissions/1503611044/
// finish times: 2
func findKthLargest(nums []int, k int) int {
	p, q := 0, len(nums)-1
	pur := len(nums) - k
	// use binary not use recursive!
	for p < q {
		pos := quickSort(nums, p, q)
		if pos == pur {
			return nums[pos]
		} else if pos < pur {
			p = pos + 1
		} else {
			q = pos - 1
		}
	}
	return nums[pur]
	// dfs
	// return findKthLargestByRange(nums, k ,p, q)
}

// func findKthLargestByRange(nums []int, k, p, q int) int {
//     pos := quickSort(nums, p, q)
//     pur := len(nums) - k
//     if pos == pur {
//         return nums[pos]
//     } else if pos < pur {
//         return findKthLargestByRange(nums, k, pos + 1, q)
//     } else {
//         return findKthLargestByRange(nums, k, p, pos - 1)
//     }
// }

func quickSort(nums []int, p, q int) int {
	randIdx := p + rand.Intn(q-p+1)
	tmp := nums[randIdx]
	nums[randIdx], nums[p] = nums[p], nums[randIdx]

	for p < q {
		for p < q && nums[q] >= tmp {
			q--
		}
		nums[p] = nums[q]
		for p < q && nums[p] <= tmp {
			p++
		}
		nums[q] = nums[p]
	}
	nums[p] = tmp
	return p
}
