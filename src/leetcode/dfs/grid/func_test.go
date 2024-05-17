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
}
