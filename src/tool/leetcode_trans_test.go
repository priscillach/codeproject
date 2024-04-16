package tool

import (
	"fmt"
	"strings"
	"testing"
)

func TestTransLCP(t *testing.T) {
	transCode := `
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		var nums1, nums2 []int
		num1, num2 := nums[i], nums[j]
		for num1 != 0 || num2 != 0 {
			if num1 > 0 {
				nums1 = append(nums1, num1%10)
			}
			if num2 > 0 {
				nums2 = append(nums2, num2%10)
			}
			num1 /= 10
			num2 /= 10
		}
		idx1 := len(nums1) - 1
		idx2 := len(nums2) - 1
		for cnt := 0; cnt < utils.Lcm(len(nums1), len(nums2)); cnt++ {
			if idx1 == 0 {
				idx1 = len(nums1) - 1
			}
			if idx2 == 0 {
				idx2 = len(nums2) - 1
			}
			num1, num2 = nums1[idx1], nums2[idx2]
			idx1--
			idx2--
			if num1 == num2 {
				continue
			} else if num1 > num2 {
				return true
			} else {
				return false
			}
		}
		return true
	})
	var res strings.Builder
	for _, num := range nums {
		res.WriteString(strconv.Itoa(num))
	}
	return res.String()
}

`
	transCode = strings.ReplaceAll(transCode, "mytreenode.", "")
	transCode = strings.ReplaceAll(transCode, "mylinkednode.", "")
	if strings.Contains(transCode, "utils.Max") {
		transCode = strings.ReplaceAll(transCode, "utils.Max", "max")
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
	if strings.Contains(transCode, "utils.Min") {
		transCode = strings.ReplaceAll(transCode, "utils.Min", "min")
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
	if strings.Contains(transCode, "utils.Lcm") {
		transCode = strings.ReplaceAll(transCode, "utils.Lcm", "lcm")
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

	if strings.Contains(transCode, "utils.Gcd") {
		transCode = strings.ReplaceAll(transCode, "utils.Gcd", "gcd")
		transCode += `
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
`
	}

	if strings.Contains(transCode, "utils.Abs") {
		transCode = strings.ReplaceAll(transCode, "utils.Abs", "abs")
		transCode += `
func Abs(num int) int {
	return int(math.Abs(float64(num)))
}
`
	}
	fmt.Println(transCode)
}
