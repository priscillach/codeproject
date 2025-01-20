package greedy

// https://leetcode.com/problems/first-completely-painted-row-or-column/description
func firstCompleteIndex(arr []int, mat [][]int) int {
	m := len(mat)
	n := len(mat[0])
	rows := make([]int, m)
	cols := make([]int, n)
	for i := range rows {
		rows[i] = n
	}
	for i := range cols {
		cols[i] = m
	}

	indexMap := make(map[int][]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			indexMap[mat[i][j]] = []int{i, j}
		}
	}

	for i := 0; i < len(arr); i++ {
		ids := indexMap[arr[i]]
		row := ids[0]
		col := ids[1]
		rows[row]--
		cols[col]--
		if rows[row] == 0 || cols[col] == 0 {
			return i
		}
	}
	return len(arr) - 1
}
