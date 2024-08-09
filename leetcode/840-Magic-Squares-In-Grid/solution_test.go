package leetcode

import "testing"

func TestNumMagicSquaresInside(t *testing.T) {
	testCases := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name:     "Example 1",
			grid:     [][]int{{4, 3, 8, 4}, {9, 5, 1, 9}, {2, 7, 6, 2}},
			expected: 1,
		},
		{
			name:     "Example 2",
			grid:     [][]int{{8}},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := numMagicSquaresInside(tc.grid)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
