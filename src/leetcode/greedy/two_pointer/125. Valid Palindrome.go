package two_pointer

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if !isValidAlphanumeric(s[left]) {
			left++
			continue
		}
		if !isValidAlphanumeric(s[right]) {
			right--
			continue
		}
		if toLowercase(s[left]) != toLowercase(s[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

func toLowercase(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + 32
	}
	return b
}

func isValidAlphanumeric(b byte) bool {
	return b >= '0' && b <= '9' || b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z'
}
