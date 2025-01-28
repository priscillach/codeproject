package string_resolve

import (
	"leetcode/src/utils/mathhelper"
	"strconv"
	"strings"
)

// https://leetcode.com/problems/restore-ip-addresses/description/
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

func restoreIpAddressesV2(s string) []string {
	IpAddresses = make([]string, 0)
	dfsRestoreIpAddressesV2(s, 0, []string{}, 0)
	return IpAddresses
}

func dfsRestoreIpAddressesV2(s string, cur int, ip []string, n int) {
	if n == 4 && cur == len(s) {
		IpAddresses = append(IpAddresses, strings.Join(ip, "."))
		return
	}
	for i := cur; i < len(s) && i < cur+3; i++ {
		if single, ok := isValidV2(s[cur : i+1]); ok {
			dfsRestoreIpAddressesV2(s, i+1, append(ip, single), n+1)
		}
	}
}

func isValidV2(s string) (string, bool) {
	if len(s) > 1 && s[0] == '0' {
		return "", false
	}
	var num int
	for i := range s {
		if s[i] < '0' || s[i] > '9' {
			return "", false
		}
		num = 10*num + int(s[i]-'0')
	}
	if num < 0 || num > 255 {
		return "", false
	}
	return strconv.Itoa(num), true
}
