package grid

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	directions := [][]int{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}
	used := make([][]bool, m)
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}
	var res []int
	curDic := 0
	x, y := 0, 0
	for i := 0; i < m*n; i++ {
		res = append(res, matrix[x][y])
		used[x][y] = true
		nX, nY := x+directions[curDic][0], y+directions[curDic][1]
		if nX >= m || nY >= n || nX < 0 || nY < 0 || used[nX][nY] {
			curDic = (curDic + 1) % 4
		}
		x, y = x+directions[curDic][0], y+directions[curDic][1]
	}
	return res
}
