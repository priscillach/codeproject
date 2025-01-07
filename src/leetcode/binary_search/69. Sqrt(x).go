package binarysearch

func mySqrt(x int) int {
	left, right := 0, x
	for left < right {
		// 要求是mid^2小于x的最小整数， 选择
		mid := left + (right-left+1)>>1
		if mid*mid <= x {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}
