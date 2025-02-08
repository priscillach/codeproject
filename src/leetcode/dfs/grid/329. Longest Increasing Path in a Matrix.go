package grid

import (
	"container/list"
	"leetcode/src/utils/mathhelper"
)

func longestIncreasingPath(matrix [][]int) int {
	maxPath := 1
	directions := [][]int{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}
	cache := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		cache[i] = make([]int, len(matrix[0]))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			maxPath = mathhelper.Max(maxPath, longestIncreasingPathDfs(matrix, cache, directions, i, j))
		}
	}
	return maxPath
}

func longestIncreasingPathDfs(matrix [][]int, cache [][]int, directions [][]int, i, j int) int {
	if cache[i][j] != 0 {
		return cache[i][j]
	}
	maxP := 1
	for _, d := range directions {
		nextI, nextJ := i+d[0], j+d[1]
		if nextI < 0 || nextI >= len(matrix) || nextJ < 0 || nextJ >= len(matrix[0]) || matrix[i][j] >= matrix[nextI][nextJ] {
			continue
		}

		maxP = mathhelper.Max(maxP, longestIncreasingPathDfs(matrix, cache, directions, nextI, nextJ)+1)
	}
	cache[i][j] = maxP
	return maxP
}

// timeout
func longestIncreasingPathV2(matrix [][]int) int {
	maxPath := 1
	directions := [][]int{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			maxPath = mathhelper.Max(maxPath, longestIncreasingPathBfs(matrix, directions, i, j))
		}
	}
	return maxPath
}

func longestIncreasingPathBfs(matrix [][]int, directions [][]int, i, j int) int {
	var queue [][]int
	cnt := 0
	queue = append(queue, []int{i, j})
	for len(queue) > 0 {
		size := len(queue)
		cnt++
		for k := 0; k < size; k++ {
			first := queue[0]
			queue = queue[1:]
			for _, d := range directions {
				nextI, nextJ := first[0]+d[0], first[1]+d[1]
				if nextI < 0 || nextI >= len(matrix) || nextJ < 0 || nextJ >= len(matrix[0]) || matrix[first[0]][first[1]] >= matrix[nextI][nextJ] {
					continue
				}
				queue = append(queue, []int{nextI, nextJ})
			}
		}
	}
	return cnt
}

// topology
func longestIncreasingPathV3(matrix [][]int) int {
	row := len(matrix)
	if row == 0 {
		return 0
	}
	col := len(matrix[0])
	if col == 0 {
		return 0
	}
	if row == 1 && col == 1 {
		return 1
	}

	dirs := [][]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	// Calculate out-degrees
	outDegrees := make([][]int, row)
	for i := range outDegrees {
		outDegrees[i] = make([]int, col)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			for _, dir := range dirs {
				nextX := i + dir[0]
				nextY := j + dir[1]

				if nextX < 0 || nextX >= row || nextY < 0 || nextY >= col || matrix[nextX][nextY] <= matrix[i][j] {
					continue
				}
				outDegrees[i][j]++
			}
		}
	}

	// Initialize queue for BFS with nodes of out-degree 0
	queue := list.New()
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if outDegrees[i][j] == 0 {
				queue.PushBack([]int{i, j})
			}
		}
	}

	cnt := 0
	for queue.Len() > 0 {
		cnt++
		size := queue.Len()
		for i := 0; i < size; i++ {
			curPos := queue.Remove(queue.Front()).([]int)
			x, y := curPos[0], curPos[1]

			for _, dir := range dirs {
				prevX := x + dir[0]
				prevY := y + dir[1]

				if prevX < 0 || prevX >= row || prevY < 0 || prevY >= col {
					continue
				}

				if matrix[prevX][prevY] >= matrix[x][y] {
					continue
				}

				outDegrees[prevX][prevY]--
				// 只有等于0会加入，多次遍历到prev，outDegrees变为负
				if outDegrees[prevX][prevY] == 0 {
					queue.PushBack([]int{prevX, prevY})
				}
			}
		}
	}

	return cnt
}
