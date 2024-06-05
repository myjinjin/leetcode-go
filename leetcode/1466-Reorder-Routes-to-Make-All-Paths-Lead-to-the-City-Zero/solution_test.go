package leetcode

import (
	"testing"
)

func TestMinReorder(t *testing.T) {
	testCases := []struct {
		name        string
		n           int
		connections [][]int
		expected    int
	}{
		{
			name:        "Example 1",
			n:           6,
			connections: [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}},
			expected:    3,
		},
		{
			name:        "Example 2",
			n:           5,
			connections: [][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}},
			expected:    2,
		},
		{
			name:        "Example 3",
			n:           3,
			connections: [][]int{{1, 0}, {2, 0}},
			expected:    0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := minReorder(tc.n, tc.connections)
			if result != tc.expected {
				t.Errorf("Expected: %d, Got: %d", tc.expected, result)
			}
		})
	}
}
