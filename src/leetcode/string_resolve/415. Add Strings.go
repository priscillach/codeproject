package string_resolve

import (
	"leetcode/src/utils/mathhelper"
	"leetcode/src/utils/stringhelper"
)

func AddStrings(num1 string, num2 string) string {
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
