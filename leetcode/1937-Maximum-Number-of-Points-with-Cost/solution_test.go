package leetcode

import "testing"

func TestMaxPoints(t *testing.T) {
	testCases := []struct {
		name     string
		points   [][]int
		expected int64
	}{
		{
			name:     "Example 1",
			points:   [][]int{{1, 2, 3}, {1, 5, 1}, {3, 1, 1}},
			expected: 9,
		},
		{
			name:     "Example 2",
			points:   [][]int{{1, 5}, {2, 3}, {4, 2}},
			expected: 11,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maxPoints(tc.points)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
