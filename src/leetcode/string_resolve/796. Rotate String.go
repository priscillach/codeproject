package string_resolve

import "strings"

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	ss := s + s
	return strings.Contains(ss, goal)
}
