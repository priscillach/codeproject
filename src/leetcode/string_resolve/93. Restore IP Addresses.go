package string_resolve

import (
	"leetcode/src/utils/mathhelper"
	"strconv"
	"strings"
)

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
	for i := 1; i <= mathhelper.Min(len(s), 3); i++ {
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
