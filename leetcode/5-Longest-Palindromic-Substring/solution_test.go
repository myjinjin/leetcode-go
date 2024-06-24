package leetcode

import "testing"

func TestLongestPalindrome(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected string
	}{
		{
			name:     "Example 1",
			s:        "babad",
			expected: "bab",
		},
		{
			name:     "Example 2",
			s:        "cbbd",
			expected: "bb",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := longestPalindrome(tc.s)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
