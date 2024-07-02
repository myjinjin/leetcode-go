package leetcode

import "testing"

func TestMaxArea(t *testing.T) {
	testCases := []struct {
		name     string
		height   []int
		expected int
	}{
		{
			name:     "Example 1",
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			name:     "Example 2",
			height:   []int{1, 1},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maxArea(tc.height)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
