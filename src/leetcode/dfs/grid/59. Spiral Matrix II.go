package grid

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	directions := [][]int{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}
	row, col := 0, 0
	d := 0
	for i := 1; i <= n*n; i++ {
		matrix[row][col] = i
		for j := 0; j < 4; j++ {
			nextR, nextC := row+directions[d][0], col+directions[d][1]
			if nextR < 0 || nextR >= n || nextC < 0 || nextC >= n || matrix[nextR][nextC] != 0 {
				d++
				d = d % 4
				continue
			} else {
				row, col = nextR, nextC
				break
			}
		}
	}
	return matrix
}
