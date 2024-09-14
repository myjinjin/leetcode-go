package leetcode

import "testing"

func TestLongestSubarray(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3, 3, 2, 2},
			expected: 2,
		},
		{
			name:     "Example 2",
			nums:     []int{1, 2, 3, 4},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := longestSubarray(tc.nums)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
