package _heap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScheduleCourse(t *testing.T) {
	assert.Equal(t, 2, scheduleCourse([][]int{{5, 5}, {4, 6}, {2, 6}}))
}
