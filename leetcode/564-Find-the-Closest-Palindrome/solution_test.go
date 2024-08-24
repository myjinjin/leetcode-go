package leetcode

import "testing"

func TestNearestPalindromic(t *testing.T) {
	testCases := []struct {
		name     string
		n        string
		expected string
	}{
		{
			name:     "Example 1",
			n:        "123",
			expected: "121",
		},
		{
			name:     "Example 2",
			n:        "1",
			expected: "0",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := nearestPalindromic(tc.n)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
