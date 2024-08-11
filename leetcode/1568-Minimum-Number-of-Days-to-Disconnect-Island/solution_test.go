package leetcode

import "testing"

func TestMinDays(t *testing.T) {
	testCases := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name:     "Example 1",
			grid:     [][]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}},
			expected: 2,
		},
		{
			name:     "Example 2",
			grid:     [][]int{{1, 1}},
			expected: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := minDays(tc.grid)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
