package leetcode

import "testing"

func TestMaxProduct(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{2, 3, -2, 4},
			expected: 6,
		},
		{
			name:     "Example 2",
			nums:     []int{-2, 0, -1},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maxProduct(tc.nums)
			if result != tc.expected {
				t.Errorf("Expected: %v, but gut %v", tc.expected, result)
			}
		})
	}
}
