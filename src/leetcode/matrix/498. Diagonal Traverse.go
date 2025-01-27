package matrix

// https://leetcode.com/problems/diagonal-traverse/description/
func findDiagonalOrder(mat [][]int) []int {
	x, y := 0, 0
	m := len(mat)
	n := len(mat[0])
	var res []int
	var topRight = true
	var isVerticalOrHorizontal = false
	for x < m && y < n {
		res = append(res, mat[x][y])
		if m == 1 {
			y++
			continue
		}
		if n == 1 {
			x++
			continue
		}
		if !isVerticalOrHorizontal && (x == 0 || x == m-1) && y < n-1 {
			isVerticalOrHorizontal = true
			topRight = !topRight
			y++
			continue
		}
		if !isVerticalOrHorizontal && (y == 0 || y == n-1) && x < m-1 {
			isVerticalOrHorizontal = true
			topRight = !topRight
			x++
			continue
		}
		if topRight {
			x--
			y++
		} else {
			y--
			x++
		}
		if isVerticalOrHorizontal {
			isVerticalOrHorizontal = false
		}
	}
	return res
}

func findDiagonalOrderV2(mat [][]int) []int {
	m := len(mat)
	n := len(mat[0])
	var res []int
	for k := 0; k < n+m; k++ {
		var list []int
		var i, j int
		if k < n {
			i, j = 0, k
		} else {
			i, j = k-n+1, n-1
		}
		for i < m && j >= 0 {
			list = append(list, mat[i][j])
			i++
			j--
		}
		if k%2 == 0 {
			// 不能使用sort.Reverse(sort.IntSlice(list))
			// sort.Reverse需要传入排序后的列表
			reverse(list)
			reverse(list)
		}
		res = append(res, list...)
	}
	return res
}

func reverse(list []int) {
	for i := 0; i < len(list)/2; i++ {
		list[i], list[len(list)-i-1] = list[len(list)-i-1], list[i]
	}
}
