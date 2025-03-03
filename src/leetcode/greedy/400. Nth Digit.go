package greedy

func findNthDigit(n int) int {
	digitNum := 9
	digitCnt := 1
	// 1 - 9     1 digit 9
	// 10 - 99   2 digit 90
	// 100 - 999 3 digit 900
	for {
		group := digitCnt * digitNum
		if n <= group {
			break
		}
		n -= group
		digitCnt++
		digitNum *= 10
	}

	start := digitNum/9 + (n-1)/digitCnt
	num := start % 10
	for i := 0; i < digitCnt-((n-1)%digitCnt+1); i++ {
		start /= 10
		num = start % 10
	}
	return num
}
