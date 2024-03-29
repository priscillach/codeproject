package tool

import (
	"fmt"
	"strings"
	"testing"
)

func TestTransLCP(t *testing.T) {
	transCode := `
var IpAddresses []string

func restoreIpAddresses(s string) []string {
	IpAddresses = make([]string, 0)
	var parts []string
	dfsRestoreIpAddresses(s, parts)
	return IpAddresses
}

func dfsRestoreIpAddresses(s string, parts []string) {
	if len(parts) == 4 {
		IpAddresses = append(IpAddresses, strings.Join(parts, "."))
		return
	}
	if len(parts) == 3 {
		if isValid(s) {
			dfsRestoreIpAddresses("", append(parts, s))
		}
		return
	}
	for i := 1; i <= utils.Min(len(s), 3); i++ {
		if isValid(s[:i]) {
			parts = append(parts, s[:i])
			dfsRestoreIpAddresses(s[i:], parts)
			parts = parts[:len(parts)-1]
		}
	}
}

func isValid(s string) bool {
	if s == "" {
			return false
	}
	if s[0] == '0' && len(s) > 1 {
		return false
	}
	t, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return false
	}
	if int(t) < 0 || int(t) > 255 {
		return false
	}
	return true
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
	fmt.Println(transCode)
}
