package matrix

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			x, y := i, j
			tmp := matrix[x][y]
			for k := 0; k < 3; k++ {
				matrix[x][y] = matrix[n-y-1][x]
				x, y = n-y-1, x
			}
			matrix[j][n-i-1] = tmp
		}
	}
}
func rotateV2(matrix [][]int) {
	n := len(matrix)
	// Transposing the matrix
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if i != j {
				matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			}
		}
	}
	// Reversing the cols of the matrix
	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			matrix[j][i], matrix[j][n-i-1] = matrix[j][n-i-1], matrix[j][i]
		}
	}
}

func rotateV3(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j <= i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix)/2; j++ {
			matrix[i][j], matrix[i][len(matrix)-j-1] = matrix[i][len(matrix)-j-1], matrix[i][j]
		}
	}
}
