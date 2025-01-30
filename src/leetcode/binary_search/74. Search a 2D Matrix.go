package binarysearch

// https://leetcode.com/problems/search-a-2d-matrix/description/
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	left, right := 0, m*n-1
	for left < right {
		mid := left + (right-left+1)>>1
		row := mid / n
		col := mid % n
		if matrix[row][col] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}
	row := left / n
	col := left % n
	return matrix[row][col] == target
}
