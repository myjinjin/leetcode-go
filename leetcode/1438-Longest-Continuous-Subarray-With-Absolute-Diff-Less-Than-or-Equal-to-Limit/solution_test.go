package leetcode

import "testing"

func TestLongestSubarray(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		limit    int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{8, 2, 4, 7},
			limit:    4,
			expected: 2,
		},
		{
			name:     "Example 2",
			nums:     []int{10, 1, 2, 4, 7, 2},
			limit:    5,
			expected: 4,
		},
		{
			name:     "Example 3",
			nums:     []int{4, 2, 2, 2, 4, 4, 2, 2},
			limit:    0,
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := longestSubarray(tc.nums, tc.limit)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
