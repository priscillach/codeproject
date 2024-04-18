package stack

func isValid(s string) bool {
	// isValidV2 chang use 3 cnt, left + 1, right - 1
	var stack []byte
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, s[i])
			continue
		}
		if stack[len(stack)-1] == '(' && s[i] == ')' {
			stack = stack[:len(stack)-1]
			continue
		}
		if stack[len(stack)-1] == '[' && s[i] == ']' {
			stack = stack[:len(stack)-1]
			continue
		}
		if stack[len(stack)-1] == '{' && s[i] == '}' {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, s[i])
	}
	return len(stack) == 0
}
