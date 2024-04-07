package string_resolve

func reverseWords(s string) string {
	var words []string
	var word []byte
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if len(word) > 0 {
				words = append(words, string(word))
				word = []byte{}
			}
			continue
		}
		if (i == 0 || s[i-1] == ' ') && s[i] != ' ' {
			word = []byte{s[i]}
		} else {
			word = append(word, s[i])
		}
	}
	if len(word) > 0 {
		words = append(words, string(word))
	}
	var res string
	for i := len(words) - 1; i >= 0; i-- {
		if i == len(words)-1 {
			res = words[i]
		} else {
			res = res + " " + words[i]
		}
	}
	return res
}
