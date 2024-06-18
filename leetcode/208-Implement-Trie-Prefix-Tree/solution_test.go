package leetcode

import "testing"

func TestSolution(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		trie := Constructor()
		trie.Insert("apple")
		if !trie.Search("apple") {
			t.Error("Search 'apple' should return true")
		}
		if trie.Search("app") {
			t.Error("Search 'app' should return false")
		}
		if !trie.StartsWith("app") {
			t.Error("StartsWith 'app' should return true")
		}
		trie.Insert("app")
		if !trie.Search("app") {
			t.Error("Search 'app' should return true")
		}
	})
}
