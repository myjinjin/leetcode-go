package leetcode

import "testing"

func TestMaxSideLength(t *testing.T) {
	testCases := []struct {
		name      string
		mat       [][]int
		threshold int
		expected  int
	}{
		{
			name:      "Example 1",
			mat:       [][]int{{1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}},
			threshold: 4,
			expected:  2,
		},
		{
			name:      "Example 2",
			mat:       [][]int{{2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}},
			threshold: 1,
			expected:  0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maxSideLength(tc.mat, tc.threshold)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
