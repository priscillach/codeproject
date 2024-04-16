package _sort

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		a, b := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
		return a+b > b+a
	})
	var res strings.Builder
	for len(nums) > 1 && nums[0] == 0 {
		nums = nums[1:]
	}
	for _, num := range nums {
		res.WriteString(strconv.Itoa(num))
	}
	return res.String()
}
