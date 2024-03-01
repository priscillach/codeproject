package leetcode

import "math/rand"

func sort(nums []int) {
	sortDfs(nums, 0, len(nums)-1)
}

func sortDfs(nums []int, p, q int) {
	pos := quickSort(nums, p, q)
	if pos+1 < q {
		sortDfs(nums, pos+1, q)

	}
	if p < pos-1 {
		sortDfs(nums, p, pos-1)
	}
}

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
