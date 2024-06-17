package leetcode

import "testing"

func TestJudgeSquareSum(t *testing.T) {
	testCases := []struct {
		name     string
		c        int
		expected bool
	}{
		{
			name:     "Example 1",
			c:        5,
			expected: true,
		},
		{
			name:     "Example 2",
			c:        3,
			expected: false,
		},
		{
			name:     "Example 3",
			c:        2,
			expected: true,
		},
		{
			name:     "Example 4",
			c:        1000,
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := judgeSquareSum(tc.c)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
