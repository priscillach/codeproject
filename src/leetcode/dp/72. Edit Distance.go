package dp

import (
	"leetcode/src/utils/mathhelper"
)

// https://leetcode.com/problems/edit-distance/submissions/1501558788/
// finish times: 2
func minDistance(word1 string, word2 string) int {
	len1 := len(word1)
	len2 := len(word2)
	distance := make([][]int, len1+1)
	for idx := range distance {
		distance[idx] = make([]int, len2+1)
		distance[idx][0] = idx
	}

	for j := 1; j <= len2; j++ {
		distance[0][j] = j
	}

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			minD := mathhelper.Min(distance[i][j-1], distance[i-1][j])
			if word1[i-1] == word2[j-1] {
				distance[i][j] = mathhelper.Min(distance[i-1][j-1], minD+1)
			} else {
				distance[i][j] = mathhelper.Min(distance[i-1][j-1], minD) + 1
			}
		}
	}
	return distance[len1][len2]
}
