package bit

// https://leetcode.com/problems/subsets/description/
// finish times: 2
func subsets(nums []int) [][]int {
	var res [][]int
	for i := 0; i < 1<<len(nums); i++ {
		set := make([]int, 0)
		for j := 0; j < len(nums); j++ {
			if i>>j&1 == 1 {
				set = append(set, nums[j])
			}
		}
		res = append(res, set)
	}
	return res
}
