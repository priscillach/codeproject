package stringhelper

import (
	"fmt"
	"strconv"
	"strings"
)

func NumByte2Int(b byte) int {
	return int(b - '0')
}

func Int2NumByte(i int) byte {
	return byte(i + '0')
}

func Int2String(num int) string {
	var res []byte
	if num == 0 {
		return "0"
	}
	for num != 0 {
		res = append([]byte{Int2NumByte(num % 10)}, res...)
		num = num / 10
	}
	return string(res)
}

func IntArr2String(nums []int) string {
	res := ""
	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 {
			res += strconv.Itoa(nums[i])
			continue
		}
		res += strconv.Itoa(nums[i]) + ","
	}
	return res
}

func String2IntArr(str string) []int {
	var nums []int
	if str == "" {
		return nums
	}
	strs := strings.Split(str, ",")
	for i := 0; i < len(strs); i++ {
		num, err := strconv.Atoi(strs[i])
		if err != nil {
			panic(fmt.Sprintf("String2IntArr err: %v", err))
		}
		nums = append(nums, num)
	}
	return nums
}

func ReverseString(str string) string {
	if len(str) == 0 {
		return str
	}
	runes := []rune(str)
	lastRune := runes[len(runes)-1]
	runes = runes[:len(runes)-1]
	return string(append([]rune{lastRune}, []rune(ReverseString(string(runes)))...))
}
