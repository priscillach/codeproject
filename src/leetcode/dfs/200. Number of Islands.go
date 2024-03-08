package dfs

//[
//["1","1","1","1","0"],
//["1","1","0","1","0"],
//["1","1","0","0","0"],
//["0","0","0","0","0"]
//]

func numIslands(grid [][]byte) int {
	num := 0
	m := len(grid)
	if m == 0 {
		return num
	}
	n := len(grid[0])
	if n == 0 {
		return num
	}
	flag := make([][]bool, m)
	for i := 0; i < m; i++ {
		flag[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !flag[i][j] {
				num++
			}
			dfsIslands(grid, i, j, flag)
		}
	}
	return num
}

func dfsIslands(grid [][]byte, x, y int, flag [][]bool) {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
		return
	}
	if grid[x][y] == '0' || flag[x][y] {
		return
	}
	flag[x][y] = true
	directions := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	for _, direction := range directions {
		nextX := x + direction[0]
		nextY := y + direction[1]
		dfsIslands(grid, nextX, nextY, flag)
	}
}
