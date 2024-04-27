package tool

import (
	"fmt"
	"strings"
	"testing"
)

func TestTransLCP(t *testing.T) {
	transCode := `
func calculate(s string) int {
	var stack []int
	var num int
	var sign byte = '+'
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + stringhelper.NumByte2Int(s[i])
		}
		if s[i] == '(' {
			// find the match ')' when cnt == 0
			var cnt int
			for j := i; j < len(s); j++ {
				if s[j] == '(' {
					cnt++
					continue
				}
				if s[j] == ')' {
					cnt--
				}
				if cnt == 0 {
					// recursively calculate the expression between the '(' and ')' i.e. i + 1 ~ j - 1
					num = calculate772(s[i+1 : j])
					// let the i move to j, next loop i++, then get the next after ')'
					i = j
					break
				}
			}
		}
		if s[i] == '+' || s[i] == '-' || i == len(s)-1 {
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			}
			sign = s[i]
			num = 0
		}
	}
	var res int
	for i := 0; i < len(stack); i++ {
		res += stack[i]
	}
	return res
}
`
	transCode = strings.ReplaceAll(transCode, "mytreenode.", "")
	transCode = strings.ReplaceAll(transCode, "mylinkednode.", "")
	if strings.Contains(transCode, "mathhelper.Max") {
		transCode = strings.ReplaceAll(transCode, "mathhelper.Max", "max")
		transCode += `
func max(nums... int) int {
	max := math.MinInt32
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
`
	}
	if strings.Contains(transCode, "mathhelper.Min") {
		transCode = strings.ReplaceAll(transCode, "mathhelper.Min", "min")
		transCode += `
func min(nums... int) int {
	min := math.MaxInt32
	for i := range nums {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}
`
	}
	if strings.Contains(transCode, "mathhelper.Lcm") {
		transCode = strings.ReplaceAll(transCode, "mathhelper.Lcm", "lcm")
		transCode += `
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
`
	}

	if strings.Contains(transCode, "mathhelper.Gcd") {
		transCode = strings.ReplaceAll(transCode, "mathhelper.Gcd", "gcd")
		transCode += `
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
`
	}

	if strings.Contains(transCode, "mathhelper.Abs") {
		transCode = strings.ReplaceAll(transCode, "mathhelper.Abs", "abs")
		transCode += `
func abs(num int) int {
	return int(math.Abs(float64(num)))
}
`
	}

	if strings.Contains(transCode, "stringhelper.NumByte2Int") {
		transCode = strings.ReplaceAll(transCode, "stringhelper.NumByte2Int", "numByte2Int")
		transCode += `
func numByte2Int(b byte) int {
	return int(b - '0')
}
`
	}
	if strings.Contains(transCode, "stringhelper.Int2NumByte") {
		transCode = strings.ReplaceAll(transCode, "stringhelper.Int2NumByte", "int2NumByte")
		transCode += `
func int2NumByte(i int) byte {
	return byte(i + '0')
}
`
	}
	if strings.Contains(transCode, "arrayhelper.FillSlice") {
		transCode = strings.ReplaceAll(transCode, "arrayhelper.FillSlice", "fillSlice")
		transCode += `
func fillSlice(nums []int, num int) {
	for i := 0; i < len(nums); i++ {
		nums[i] = num
	}
}
`
	}

	fmt.Println(transCode)
}
