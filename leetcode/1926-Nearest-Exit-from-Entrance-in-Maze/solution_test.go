package leetcode

import "testing"

func TestNearestExit(t *testing.T) {
	testCases := []struct {
		name     string
		maze     [][]byte
		entrance []int
		expected int
	}{
		{
			name:     "Example 1",
			maze:     [][]byte{{'+', '+', '.', '+'}, {'.', '.', '.', '+'}, {'+', '+', '+', '.'}},
			entrance: []int{1, 2},
			expected: 1,
		},
		{
			name:     "Example 2",
			maze:     [][]byte{{'+', '+', '+'}, {'.', '.', '.'}, {'+', '+', '+'}},
			entrance: []int{1, 0},
			expected: 2,
		},
		{
			name:     "Example 3",
			maze:     [][]byte{{'.', '+'}},
			entrance: []int{0, 0},
			expected: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := nearestExit(tc.maze, tc.entrance)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
