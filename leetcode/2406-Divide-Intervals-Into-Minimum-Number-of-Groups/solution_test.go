package leetcode

import "testing"

func TestMinGroups(t *testing.T) {
	testCases := []struct {
		name      string
		intervals [][]int
		expected  int
	}{
		{
			name:      "Example 1",
			intervals: [][]int{{5, 10}, {6, 8}, {1, 5}, {2, 3}, {1, 10}},
			expected:  3,
		},
		{
			name:      "Example 2",
			intervals: [][]int{{1, 3}, {5, 6}, {8, 10}, {11, 13}},
			expected:  1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := minGroups(tc.intervals)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
