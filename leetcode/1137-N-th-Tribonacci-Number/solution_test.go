package leetcode

import "testing"

func TestTribonacci(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "Example 1",
			n:        4,
			expected: 4,
		},
		{
			name:     "Example 2",
			n:        25,
			expected: 1389537,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tribonacci(tc.n)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
