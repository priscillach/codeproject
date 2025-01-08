package stringhelper

import (
	"fmt"
	"testing"
)

func TestInt2String(t *testing.T) {
	fmt.Println(Int2String(0))
	fmt.Println(Int2String(123))
}

func TestReverseString(t *testing.T) {
	fmt.Println(ReverseString("hello，世界"))
}
