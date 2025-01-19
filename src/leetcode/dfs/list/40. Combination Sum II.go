package list

import "sort"

var combinationSumRes40 [][]int

func combinationSum40(candidates []int, target int) [][]int {
	combinationSumRes40 = make([][]int, 0)
	sort.Ints(candidates)
	combinationSumCore40(candidates, 0, []int{}, 0, target)
	return combinationSumRes40
}

func combinationSumCore40(candidates []int, first int, chosen []int, sum, target int) {
	if sum > target {
		return
	}
	if sum == target {
		copyArr := make([]int, len(chosen))
		copy(copyArr, chosen)
		combinationSumRes = append(combinationSumRes, copyArr)
		return
	}
	for i := first; i < len(candidates); i++ {
		if i > first && candidates[i] == candidates[i-1] {
			continue
		}
		combinationSumCore40(candidates, i+1, append(chosen, candidates[i]), sum+candidates[i], target)
	}
}
