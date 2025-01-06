package string_resolve

func firstNonRepeatedCharacter(s string) byte {
	charCount := make(map[byte]int)
	for _, c := range s {
		charCount[byte(c)]++
	}
	for _, c := range s {
		if charCount[byte(c)] == 1 {
			return byte(c)
		}
	}
	return 0
}
