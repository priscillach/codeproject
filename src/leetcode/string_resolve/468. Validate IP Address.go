package string_resolve

import (
	"strconv"
	"strings"
)

func validIPAddress(queryIP string) string {
	if IsValidIPv4(queryIP) {
		return "IPv4"
	}
	if IsValidIPv6(queryIP) {
		return "IPv6"
	}
	return "Neither"
}

func IsValidIPv4(queryIP string) bool {
	parts := strings.Split(queryIP, ".")
	if len(parts) != 4 {
		return false
	}
	for _, part := range parts {
		if len(part) > 1 && part[0] == '0' {
			return false
		}
		num, err := strconv.Atoi(part)
		if err != nil {
			return false
		}
		if num < 0 || num > 255 {
			return false
		}
	}
	return true
}

func IsValidIPv6(queryIP string) bool {
	parts := strings.Split(queryIP, ":")
	if len(parts) != 8 {
		return false
	}
	for _, part := range parts {
		if len(part) > 4 || len(part) == 0 {
			return false
		}
		for i := 0; i < len(part); i++ {
			c := part[i]
			if (c > 'f' || c < 'a') && (c > 'F' || c < 'A') && (c > '9' || c < '0') {
				return false
			}
		}
	}
	return true
}
