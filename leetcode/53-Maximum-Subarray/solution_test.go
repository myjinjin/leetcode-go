package leetcode

import "testing"

func TestMaxSubArray(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: 6,
		},
		{
			name:     "Example 2",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "Example 3",
			nums:     []int{5, 4, -1, 7, 8},
			expected: 23,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maxSubArray(tc.nums)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
