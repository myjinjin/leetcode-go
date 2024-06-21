package leetcode

import "testing"

func TestFindMinArrowShots(t *testing.T) {
	testCases := []struct {
		name     string
		points   [][]int
		expected int
	}{
		{
			name:     "Example 1",
			points:   [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}},
			expected: 2,
		},
		{
			name:     "Example 2",
			points:   [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
			expected: 4,
		},
		{
			name:     "Example 3",
			points:   [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
			expected: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findMinArrowShots(tc.points)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
