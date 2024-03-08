package greedy

import "sort"

func findLeastNumOfUniqueInts(arr []int, k int) int {
	countMap := make(map[int]int)
	for _, item := range arr {
		countMap[item]++
	}
	countList := make([]int, 0)
	for _, count := range countMap {
		countList = append(countList, count)
	}
	sort.Slice(countList, func(i, j int) bool {
		return countList[i] < countList[j]
	})
	for idx, count := range countList {
		if k == 0 {
			return len(countList) - idx
		}
		if k < count {
			return len(countList) - idx
		}

		k -= count
	}
	return 0
}
