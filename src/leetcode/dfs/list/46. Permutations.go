package list

import "fmt"

var res [][]int

func permute(nums []int) [][]int {
	res = [][]int{}
	dfs(nums, []int{})
	fmt.Println(res)
	return res
}

func dfs(nums []int, permutation []int) {
	if len(permutation) == len(nums) {
		target := make([]int, len(permutation))
		copy(target, permutation)
		res = append(res, target)
		return
	}
	for idx, num := range nums {
		if num == -11 {
			continue
		}
		nums[idx] = -11
		permutation = append(permutation, num)
		dfs(nums, permutation)
		nums[idx] = num
		permutation = permutation[:len(permutation)-1]
	}
}
