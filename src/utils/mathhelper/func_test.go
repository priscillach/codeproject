package mathhelper

import (
	"fmt"
	"testing"
)

func TestSwapWithThirdVariable(t *testing.T) {
	a, b := 1, 2
	a, b = SwapWithThirdVariable(a, b)
	fmt.Println(a, b)
}
