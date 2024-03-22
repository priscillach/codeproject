package utils

func NumByte2Int(b byte) int {
	return int(b - '0')
}

func Int2NumByte(i int) byte {
	return byte(i + '0')
}
