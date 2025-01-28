package string_resolve

import (
	"leetcode/src/utils/mathhelper"
	"leetcode/src/utils/stringhelper"
	"strconv"
)

// https://leetcode.com/problems/add-strings/description/
func addStrings(num1 string, num2 string) string {
	len1 := len(num1)
	len2 := len(num2)
	lenM := mathhelper.Max(len1, len2)
	var res []byte
	var bit1, bit2, add int
	for i := 0; i < lenM; i++ {
		idx1 := len1 - i - 1
		idx2 := len2 - i - 1
		if idx1 >= 0 {
			bit1 = stringhelper.NumByte2Int(num1[idx1])
		} else {
			bit1 = 0
		}
		if idx2 >= 0 {
			bit2 = stringhelper.NumByte2Int(num2[idx2])
		} else {
			bit2 = 0
		}
		if bit1+bit2+add >= 10 {
			res = append([]byte{stringhelper.Int2NumByte((bit1 + bit2 + add) % 10)}, res...)
			add = 1
		} else {
			res = append([]byte{stringhelper.Int2NumByte(bit1 + bit2 + add)}, res...)
			add = 0
		}
	}
	if add > 0 {
		res = append([]byte{stringhelper.Int2NumByte(add)}, res...)
	}
	return string(res)
}

func addStringsV2(num1 string, num2 string) string {
	carry := 0
	var res string
	for i := 0; i < mathhelper.Max(len(num1), len(num2)); i++ {
		var a, b int
		if i < len(num1) {
			a = int(num1[len(num1)-i-1] - '0')
		}
		if i < len(num2) {
			b = int(num2[len(num2)-i-1] - '0')
		}
		sum := a + b + carry
		carry = sum / 10
		res = strconv.Itoa(sum%10) + res
	}
	if carry > 0 {
		res = strconv.Itoa(carry) + res
	}
	return res
}
