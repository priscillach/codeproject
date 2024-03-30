package string_resolve

import (
	"fmt"
	"testing"
)

func TestAddStrings(t *testing.T) {
	fmt.Println(addStrings("125", "92"))
}

func TestRestoreIpAddresses(t *testing.T) {
	fmt.Println(restoreIpAddresses("101023"))
}

func TestMyAtoi(t *testing.T) {
	myAtoi("4193 with words")
	fmt.Println('0', '9')
}
