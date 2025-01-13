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

func spiralOrderV2(matrix [][]int) []int {
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	m := len(matrix)
	n := len(matrix[0])
	used := make([][]bool, m)
	for i := range used {
		used[i] = make([]bool, n)
	}

	var res []int
	i, j, nextI, nextJ := 0, 0, 0, 0
	d := 0
	for {
		res = append(res, matrix[i][j])
		used[i][j] = true
		firstD := d
		nextI, nextJ = i+directions[d][0], j+directions[d][1]
		for needTurn(used, nextI, nextJ) {
			d++
			d %= 4
			if firstD == d {
				return res
			}
			nextI, nextJ = i+directions[d][0], j+directions[d][1]
		}
		i, j = nextI, nextJ
	}
	return res
}

func needTurn(used [][]bool, i, j int) bool {
	m := len(used)
	n := len(used[0])
	if i < 0 || i >= m || j < 0 || j >= n {
		return true
	}
	return used[i][j]
}
