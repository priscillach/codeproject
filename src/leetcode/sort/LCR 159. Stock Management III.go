package _sort

// https://leetcode.cn/problems/zui-xiao-de-kge-shu-lcof/
func inventoryManagement(stock []int, cnt int) []int {
	var res []int
	n := len(stock)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(stock, n, i)
	}
	for i := n - 1; i > n-1-cnt; i-- {
		stock[i], stock[0] = stock[0], stock[i]
		res = append(res, stock[i])
		heapify(stock, i, 0)
	}
	return res
}

func heapify(nums []int, n, i int) {
	smallest := i
	left, right := 2*i+1, 2*i+2
	if left < n && nums[smallest] > nums[left] {
		smallest = left
	}
	if right < n && nums[smallest] > nums[right] {
		smallest = right
	}

	if smallest != i {
		nums[i], nums[smallest] = nums[smallest], nums[i]
		heapify(nums, n, smallest)
	}
}
