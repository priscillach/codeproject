package grid

import (
	"fmt"
	"testing"
)

func TestNumIslands(t *testing.T) {
	grid := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}
	fmt.Println(numIslands(grid))
}

func TestWordSearch(t *testing.T) {
	grid := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'}}
	exist(grid, "ABCESEEEFS")

	existV2(grid, "ABCB")

}

func TestSpiralOrder(t *testing.T) {
	spiralOrderV2([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
}

func TestGenerateMatrix(t *testing.T) {
	generateMatrix(3)
}

func TestLongestIncreasingPath(t *testing.T) {
	longestIncreasingPath([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}})
	longestIncreasingPathV2([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}})
	longestIncreasingPathV3([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}})
}
