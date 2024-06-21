package leetcode

import "testing"

func TestEraseOverlapIntervals(t *testing.T) {
	testCases := []struct {
		name      string
		intervals [][]int
		expected  int
	}{
		{
			name:      "Example 1",
			intervals: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
			expected:  1,
		},
		{
			name:      "Example 2",
			intervals: [][]int{{1, 2}, {1, 2}, {1, 2}},
			expected:  2,
		},
		{
			name:      "Example 3",
			intervals: [][]int{{1, 2}, {2, 3}},
			expected:  0,
		},
		{
			name:      "Example 4",
			intervals: [][]int{{0, 2}, {1, 3}, {2, 4}, {3, 5}, {4, 6}},
			expected:  2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := eraseOverlapIntervals(tc.intervals)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
