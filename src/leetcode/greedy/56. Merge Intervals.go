package greedy

import (
	"leetcode/src/utils/mathhelper"
	"sort"
)

func mergeIntervals(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{intervals[0]}

	for idx := range intervals {
		if res[len(res)-1][1] < intervals[idx][0] {
			res = append(res, intervals[idx])
			continue
		}
		res[len(res)-1][1] = mathhelper.Max(res[len(res)-1][1], intervals[idx][1])
	}
	return res
}
