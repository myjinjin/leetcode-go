package leetcode

import "testing"

func TestCloseStrings(t *testing.T) {
	testCases := []struct {
		word1    string
		word2    string
		expected bool
	}{
		{"abc", "bca", true},
		{"a", "aa", false},
		{"cabbba", "abbccc", true},
	}

	for _, tc := range testCases {
		result := closeStrings(tc.word1, tc.word2)
		if result != tc.expected {
			t.Errorf("closeStrings(%q, %q) = %v, expected %v", tc.word1, tc.word2, result, tc.expected)
		}
	}
}
