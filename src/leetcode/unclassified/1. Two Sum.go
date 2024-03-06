package unclassified

func twoSum(nums []int, target int) []int {
	left := make(map[int]int)
	for idx, num := range nums {
		left[target-num] = idx
	}
	for idx, num := range nums {
		if idx2, ok := left[num]; ok && idx != idx2 {
			return []int{idx, idx2}
		}
	}
	return nil
}
