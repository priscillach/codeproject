package greedy

import "leetcode/src/utils/mathhelper"

// https://leetcode.com/problems/candy/description/
func candy(ratings []int) int {
	var sum int
	allocates := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		allocates[i] = 1
		if i > 0 && ratings[i] > ratings[i-1] {
			allocates[i] = allocates[i-1] + 1
			continue
		}
		for j := i - 1; j >= 0; j-- {
			if ratings[j] > ratings[j+1] && allocates[j] <= allocates[j+1] {
				allocates[j] = allocates[j+1] + 1
			}
		}
	}
	for _, num := range allocates {
		sum += num
	}
	return sum
}

func candyV2(ratings []int) int {
	var sum int
	allocates := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		allocates[i] = 1
		if i > 0 && ratings[i] > ratings[i-1] {
			allocates[i] = allocates[i-1] + 1
		}
	}

	for i := len(ratings) - 1; i >= 0; i-- {
		if i < len(ratings)-1 && ratings[i] > ratings[i+1] {
			allocates[i] = mathhelper.Max(allocates[i], allocates[i+1]+1)
		}
	}
	for _, num := range allocates {
		sum += num
	}
	return sum
}
