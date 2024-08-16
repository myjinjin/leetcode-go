package leetcode

import "testing"

func TestMaxDistance(t *testing.T) {
	testCases := []struct {
		name     string
		arrays   [][]int
		expected int
	}{
		{
			name:     "Example 1",
			arrays:   [][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}},
			expected: 4,
		},
		{
			name:     "Example 2",
			arrays:   [][]int{{1}, {1}},
			expected: 0,
		},
		{
			name:     "Example 3",
			arrays:   [][]int{{-10, -9, -9, -3, -1, -1, 0}, {-5}, {4}, {-8}, {-9, -6, -5, -4, -2, 2, 3}, {-3, -3, -2, -1, 0}},
			expected: 14,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maxDistance(tc.arrays)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
