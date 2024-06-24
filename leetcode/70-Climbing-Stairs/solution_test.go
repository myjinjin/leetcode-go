package leetcode

import "testing"

func TestClimbStairs(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "Example 1",
			n:        2,
			expected: 2,
		},
		{
			name:     "Example 2",
			n:        3,
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := climbStairs(tc.n)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}

}
