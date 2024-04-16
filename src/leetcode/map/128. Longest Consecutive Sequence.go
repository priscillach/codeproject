package _map

func longestConsecutive(nums []int) int {
	maxLen := 0
	set := map[int]struct{}{}
	for i := 0; i < len(nums); i++ {
		set[nums[i]] = struct{}{}
	}

	for i := 0; i < len(nums); i++ {
		if _, ok := set[nums[i]-1]; ok {
			continue
		} else {
			cnt := 1
			cur := nums[i] + 1
			for {
				if _, ok2 := set[cur]; !ok2 {
					break
				}
				cur++
				cnt++
			}
			if cnt > maxLen {
				maxLen = cnt
			}
		}
	}
	return maxLen
}
