package double_point

import "sort"

func bagOfTokensScore(tokens []int, power int) int {
	if len(tokens) == 0 {
		return 0
	}
	sort.Ints(tokens)
	left, right := 0, len(tokens)-1
	score := 0
	for left < len(tokens) && power >= tokens[left] {
		score++
		power -= tokens[left]
		left++
	}
	for left < right && score > 0 {
		score--
		power += tokens[right]
		for left < right && power >= tokens[left] {
			score++
			power -= tokens[left]
			left++
		}
		right--
	}
	return score
}
