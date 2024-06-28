package leetcode

import "testing"

func TestMaximumImportance(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		roads    [][]int
		expected int64
	}{
		{
			name:     "Example 1",
			n:        5,
			roads:    [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 2}, {1, 3}, {2, 4}},
			expected: 43,
		},
		{
			name:     "Example 2",
			n:        5,
			roads:    [][]int{{0, 3}, {2, 4}, {1, 3}},
			expected: 20,
		},
	}

	for _, tc := range testCases {
		result := maximumImportance(tc.n, tc.roads)
		if result != tc.expected {
			t.Errorf("Expected: %v, but got %v", tc.expected, result)
		}
	}
}
