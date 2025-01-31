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
	MyAtoi("4193 with words")
	fmt.Println('0', '9')
}

func TestCompareVersion(t *testing.T) {
	compareVersion("1.01", "1.001")
}

func TestMultiply(t *testing.T) {
	Multiply("123", "456")
	Multiply("9133", "0")

	MultiplyV2("9133", "0")
	MultiplyV2("123", "456")
}

func TestFirstNonRepeatedCharacter(t *testing.T) {
	firstNonRepeatedCharacter("")
}

func TestPermutations(t *testing.T) {
	permutationsRes = permutations("1123")
	fmt.Println(permutationsRes)
	fmt.Println(len(permutationsRes))
}

func TestSwapString(t *testing.T) {
	fmt.Println(swapString("Hello", "Damn"))
}
