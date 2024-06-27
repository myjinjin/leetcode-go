package leetcode

import "testing"

func TestCanPartition(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 5, 11, 5},
			expected: true,
		},
		{
			name:     "Example 2",
			nums:     []int{1, 2, 3, 5},
			expected: false,
		},
		{
			name:     "Example 3",
			nums:     []int{2, 2, 1, 1},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := canPartition(tc.nums)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
