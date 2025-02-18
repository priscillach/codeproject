package array_shuffle

import "math/rand"

// https://leetcode.com/problems/shuffle-an-array/description/
// Fisher-Yates Algorithm

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums: nums}
}

func (this *Solution) Reset() []int {
	return this.nums
}

func (this *Solution) Shuffle() []int {
	shuffle := make([]int, len(this.nums))
	copy(shuffle, this.nums)
	for i := 0; i < len(this.nums)-1; i++ {
		idx := rand.Intn(len(this.nums)-i) + i
		shuffle[i], shuffle[idx] = shuffle[idx], shuffle[i]
	}
	return shuffle
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
