package greedy

func iceBreakingGame(num int, target int) int {
	pos := 0
	for i := 2; i <= num; i++ {
		pos = (pos + target) % i
	}
	return pos
}
