package leetcode

import "testing"

func TestHasTrailingZeros(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3, 4, 5},
			expected: true,
		},
		{
			name:     "Example 2",
			nums:     []int{2, 4, 8, 16},
			expected: true,
		},
		{
			name:     "Example 3",
			nums:     []int{1, 3, 5, 7, 9},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := hasTrailingZeros(tc.nums)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
