package list

var res47 [][]int
var cnt []int

// -10 <= nums[i] <= 10
func permuteUnique(nums []int) [][]int {
	res47 = [][]int{}
	cnt = make([]int, 21)
	for _, num := range nums {
		cnt[num+10]++
	}
	dfs47(nums, []int{})
	return res47
}

func dfs47(nums []int, permutation []int) {
	if len(permutation) == len(nums) {
		target := make([]int, len(permutation))
		copy(target, permutation)
		res47 = append(res47, target)
		return
	}
	used := make([]bool, 21)
	for _, num := range nums {
		if cnt[num+10] == 0 {
			continue
		}
		if used[num+10] {
			continue
		}
		used[num+10] = true
		cnt[num+10]--
		permutation = append(permutation, num)
		dfs47(nums, permutation)
		cnt[num+10]++
		permutation = permutation[:len(permutation)-1]
	}
}
