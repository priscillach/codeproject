package grid

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if visited[i][j] || grid[i][j] == 0 {
				continue
			}
			area := maxAreaOfIslandCore(grid, visited, i, j, [][]int{
				{0, 1}, {0, -1}, {1, 0}, {-1, 0},
			})
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func maxAreaOfIslandCore(grid [][]int, visited [][]bool, i, j int, directions [][]int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 || visited[i][j] {
		return 0
	}
	visited[i][j] = true
	var sum int
	for _, direction := range directions {
		sum += maxAreaOfIslandCore(grid, visited, i+direction[0], j+direction[1], directions)
	}
	return sum + 1
}
