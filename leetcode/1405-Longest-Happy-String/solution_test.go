package leetcode

import "testing"

func TestLongestDiverseString(t *testing.T) {
	testCases := []struct {
		name     string
		a        int
		b        int
		c        int
		expected string
	}{
		{
			name:     "Example 1",
			a:        1,
			b:        1,
			c:        7,
			expected: "ccaccbcc",
		},
		{
			name:     "Example 2",
			a:        7,
			b:        1,
			c:        0,
			expected: "aabaa",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := longestDiverseString(tc.a, tc.b, tc.c)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
