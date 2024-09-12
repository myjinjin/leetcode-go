package leetcode

import "testing"

func TestCountConsistentStrings(t *testing.T) {
	testCases := []struct {
		name     string
		allowed  string
		words    []string
		expected int
	}{
		{
			name:     "Example 1",
			allowed:  "ab",
			words:    []string{"ad", "bd", "aaab", "baa", "badab"},
			expected: 2,
		},
		{
			name:     "Example 2",
			allowed:  "abc",
			words:    []string{"a", "b", "c", "ab", "ac", "bc", "abc"},
			expected: 7,
		},
		{
			name:     "Example 3",
			allowed:  "cad",
			words:    []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"},
			expected: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := countConsistentStrings(tc.allowed, tc.words)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
