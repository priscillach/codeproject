package matrix

import "testing"

func TestRotate(t *testing.T) {
	rotate([][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	})
}

func TestFindDiagonalOrder(t *testing.T) {
	findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	findDiagonalOrder([][]int{{6, 9, 7}})

	findDiagonalOrderV2([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	findDiagonalOrderV2([][]int{{2, 5}, {8, 4}, {0, -1}})
}
