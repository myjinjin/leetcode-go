package leetcode

import "testing"

func TestFindTheWinner(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		k        int
		expected int
	}{
		{
			name:     "Example 1",
			n:        5,
			k:        2,
			expected: 3,
		},
		{
			name:     "Example 2",
			n:        6,
			k:        5,
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findTheWinner(tc.n, tc.k)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
