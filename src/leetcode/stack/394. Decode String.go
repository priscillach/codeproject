package stack

import (
	"leetcode/src/utils/stringhelper"
	"strings"
)

func decodeString(s string) string {
	strStack := []string{""}
	var numStack []int
	var num int
	for i := 0; i < len(s); i++ {
		if s[i] <= 'z' && s[i] >= 'a' {
			strStack[len(strStack)-1] += string(s[i])
		} else if s[i] != '[' && s[i] != ']' {
			num = num*10 + stringhelper.NumByte2Int(s[i])
		} else if s[i] == '[' {
			numStack = append(numStack, num)
			// add a new empty str into stack as encountering a new sub "[...]"
			strStack = append(strStack, "")
			num = 0
		} else {
			str := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			curNum := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			// add the repeats to the last str in the stack
			strStack[len(strStack)-1] += strings.Repeat(str, curNum)
		}
	}
	return strStack[0]
}
