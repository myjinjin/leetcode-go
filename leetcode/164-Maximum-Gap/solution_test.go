package leetcode

import "testing"

func TestMaximumGap(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{3, 6, 9, 1},
			expected: 3,
		},
		{
			name:     "Example 2",
			nums:     []int{10},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maximumGap(tc.nums)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
