package leetcode

import "testing"

func TestMinimumCost(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected int64
	}{
		{
			name:     "Example 1",
			s:        "0011",
			expected: 2,
		},
		{
			name:     "Example 2",
			s:        "010101",
			expected: 9,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := minimumCost(tc.s)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
