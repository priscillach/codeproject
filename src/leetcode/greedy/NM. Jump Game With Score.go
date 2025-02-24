package greedy

import "leetcode/src/utils/mathhelper"

func maxScorePath(arr []int, startScore int) (bool, int, []int) {
	path := []int{0}
	arr[0] += startScore
	cur := 0
	score := arr[0]
	for cur < len(arr)-1 {
		prev := cur
		for i := cur + 1; i < mathhelper.Min(len(arr), cur+score+1); i++ {
			if arr[i] > 0 || i == len(arr)-1 {
				score -= i - cur
				score += arr[i]
				path = append(path, i)
				cur = i
				if i == len(arr)-1 {
					return true, score, path
				}
				break
			}
		}
		if prev == cur {
			return false, score, path
		}
	}
	return false, score, path
}
