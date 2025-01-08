package string_resolve

var permutationsRes []string
var runeCountMap map[rune]int

func permutations(str string) []string {
	runeCountMap = make(map[rune]int)
	for _, r := range str {
		runeCountMap[r]++
	}
	permutationsRes = []string{}
	permutationCore([]rune(str), []rune{})
	return permutationsRes
}

func permutationCore(runes []rune, permutation []rune) {
	if len(permutation) == len(runes) {
		permutationsRes = append(permutationsRes, string(permutation))
		return
	}
	used := make(map[rune]bool)
	for i := 0; i < len(runes); i++ {
		if used[runes[i]] {
			continue
		}
		if runeCountMap[runes[i]] <= 0 {
			continue
		}
		used[runes[i]] = true
		runeCountMap[runes[i]]--
		permutation = append(permutation, runes[i])
		permutationCore(runes, permutation)
		runeCountMap[runes[i]]++
		permutation = permutation[:len(permutation)-1]
	}
}
