package tool

import (
	"fmt"
	"strings"
	"testing"
)

func TestTransLCP(t *testing.T) {
	transCode := `
func canJump(nums []int) bool {
	curMaxPos := 0
	curPos := 0
	for curPos < curMaxPos && curPos < len(nums)-1 {
		curMaxPos = mathhelper.Max(curMaxPos, curPos+nums[curPos])
		curPos++
	}
	return curPos == len(nums)-1
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
