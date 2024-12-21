package topological_graph

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindOrder(t *testing.T) {
	findOrder(2, [][]int{{0, 1}})
}

func TestDetectCircularDependency(t *testing.T) {
	isCircle, path := detectCircularDependency([][]int{{0, 2}, {1, 2}, {2, 3}, {2, 4}})
	fmt.Println(path)
	assert.Equal(t, true, isCircle)
}
