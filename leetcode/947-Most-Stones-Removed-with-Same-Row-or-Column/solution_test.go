package leetcode

import "testing"

func TestRemoveStones(t *testing.T) {
	testCases := []struct {
		name     string
		stones   [][]int
		expected int
	}{
		{
			name:     "Example 1",
			stones:   [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}},
			expected: 5,
		},
		{
			name:     "Example 2",
			stones:   [][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}},
			expected: 3,
		},
		{
			name:     "Example 3",
			stones:   [][]int{{0, 0}},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := removeStones(tc.stones)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
