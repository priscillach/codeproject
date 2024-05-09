package matrix

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
