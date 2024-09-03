package leetcode

import "testing"

func TestLongestNiceSubarray(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 3, 8, 48, 10},
			expected: 3,
		},
		{
			name:     "Example 2",
			nums:     []int{3, 1, 5, 11, 13},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := longestNiceSubarray(tc.nums)
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})
	}
}
