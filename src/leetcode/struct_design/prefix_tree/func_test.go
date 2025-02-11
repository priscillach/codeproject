package prefix_tree

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := Constructor()
	trie.Insert("hotdog")
	fmt.Println(trie.StartsWith("dog"))
}
