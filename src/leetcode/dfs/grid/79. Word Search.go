package grid

// don't need map[i-j-idx] = false
// special case
// ABCE
// SFES
// ADEE
// ABCESEEEFS
func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if existCore(board, i, j, 0, word, visited, directions) {
				return true
			}
		}
	}
	return false
}

func existCore(board [][]byte, i, j, idx int, word string, visited [][]bool, directions [][]int) bool {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || visited[i][j] {
		return false
	}
	if idx == len(word)-1 && board[i][j] == word[idx] {
		return true
	}
	if board[i][j] != word[idx] {
		return false
	}
	visited[i][j] = true
	for _, direction := range directions {
		if existCore(board, i+direction[0], j+direction[1], idx+1, word, visited, directions) {
			return true
		}
	}
	visited[i][j] = false
	return false
}
