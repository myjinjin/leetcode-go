package leetcode

import "testing"

func TestRangeSum(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		n        int
		left     int
		right    int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3, 4},
			n:        4,
			left:     1,
			right:    5,
			expected: 13,
		},
		{
			name:     "Example 2",
			nums:     []int{1, 2, 3, 4},
			n:        4,
			left:     3,
			right:    4,
			expected: 6,
		},
		{
			name:     "Example 3",
			nums:     []int{1, 2, 3, 4},
			n:        4,
			left:     1,
			right:    10,
			expected: 50,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := rangeSum(tc.nums, tc.n, tc.left, tc.right)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
