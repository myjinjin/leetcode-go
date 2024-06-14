package leetcode

import "testing"

func TestUniquePaths(t *testing.T) {
	testCases := []struct {
		name     string
		m        int
		n        int
		expected int
	}{
		{
			name:     "Example 1",
			m:        3,
			n:        7,
			expected: 28,
		},
		{
			name:     "Example 2",
			m:        3,
			n:        2,
			expected: 3,
		},
		{
			name:     "Example 3",
			m:        23,
			n:        12,
			expected: 193536720,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := uniquePaths(tc.m, tc.n)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
