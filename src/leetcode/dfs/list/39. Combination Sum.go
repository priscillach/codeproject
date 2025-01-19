package list

var combinationSumRes [][]int

// https://leetcode.com/problems/combination-sum/
func combinationSum(candidates []int, target int) [][]int {
	combinationSumRes = make([][]int, 0)
	combinationSumCore(candidates, []int{}, 0, 0, target)
	return combinationSumRes
}

func combinationSumCore(candidates []int, chosen []int, cur, sum, target int) {
	if cur >= len(candidates) {
		return
	}
	if sum > target {
		return
	}
	if sum == target {
		copyArr := make([]int, len(chosen))
		copy(copyArr, chosen)
		combinationSumRes = append(combinationSumRes, copyArr)
		return
	}
	combinationSumCore(candidates, append(chosen, candidates[cur]), cur, sum+candidates[cur], target)
	combinationSumCore(candidates, chosen, cur+1, sum, target)
}
