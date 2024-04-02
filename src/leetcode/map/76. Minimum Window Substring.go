package _map

func minWindow(s string, t string) string {
	cntDic := make([]int, 52)
	coverFlag := 0
	subStr := ""
	for i := 0; i < len(t); i++ {
		idx := findIdx(t[i])
		// let letter cnt > 0 in t
		if cntDic[idx] == 0 {
			coverFlag++
		}
		cntDic[idx]++
	}

	left := 0
	for right := 0; right < len(s); right++ {
		var flag bool
		idx := findIdx(s[right])
		// letter cnt not in t will <= 0
		cntDic[idx]--
		if cntDic[idx] == 0 {
			coverFlag--
		}
		if coverFlag == 0 {
			flag = true
		}
		for coverFlag == 0 && left <= right {
			idx = findIdx(s[left])
			if cntDic[idx] == 0 {
				coverFlag++
			}
			cntDic[idx]++
			left++
		}
		if flag && subStr == "" || (subStr != "" && right-left+2 < len(subStr)) {
			subStr = s[left-1 : right+1]
		}
	}
	return subStr
}

func findIdx(b byte) int {
	if b >= 'a' && b <= 'z' {
		return int(b - 'a')
	}
	if b >= 'A' && b <= 'Z' {
		return int(b-'A') + 26
	}
	return -1
}
