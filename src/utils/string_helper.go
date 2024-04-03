package utils

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
