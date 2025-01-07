package mathhelper

import "math"

func Max(nums ...int) int {
	maxNum := math.MinInt32
	for i := range nums {
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
	}
	return maxNum
}

func Min(nums ...int) int {
	minNum := math.MaxInt32
	for i := range nums {
		if nums[i] < minNum {
			minNum = nums[i]
		}
	}
	return minNum
}

func Abs(num int) int {
	return int(math.Abs(float64(num)))
}

// 最大公因数
func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// 最小公倍数
func Lcm(a, b int) int {
	return a * b / Gcd(a, b)
}

func LowBit(num int) int {
	return num & -num
}

func SwapWithThirdVariable(a, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b
}
