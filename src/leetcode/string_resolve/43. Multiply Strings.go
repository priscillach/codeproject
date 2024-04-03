package string_resolve

import "leetcode/src/utils"

func Multiply(num1 string, num2 string) string {
	var res = "0"
	for i := 0; i < len(num2); i++ {
		idx2 := len(num2) - i - 1
		a := utils.NumByte2Int(num2[idx2])
		for j := 0; j < len(num1); j++ {
			idx1 := len(num1) - j - 1
			b := utils.NumByte2Int(num1[idx1])
			add := addZero(utils.Int2String(a*b), j+i)
			res = AddStrings(res, add)
		}
	}
	return res
}

func addZero(num string, zeros int) string {
	if num == "0" {
		return num
	}
	for i := 0; i < zeros; i++ {
		num = num + "0"
	}
	return num
}

func MultiplyV2(num1 string, num2 string) string {
	// the max length of multipliers is the sum of the two digits
	res := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			tens, ones := i+j, i+j+1
			mul := utils.NumByte2Int(num1[i])*utils.NumByte2Int(num2[j]) + res[ones]
			res[ones] = mul % 10
			// if res[tens] >= 10, will be mod in next iteration
			res[tens] += mul / 10
		}
	}
	for len(res) > 1 && res[0] == 0 {
		res = res[1:]
	}
	var product []byte
	for i := 0; i < len(res); i++ {
		product = append(product, utils.Int2NumByte(res[i]))
	}
	return string(product)
}
