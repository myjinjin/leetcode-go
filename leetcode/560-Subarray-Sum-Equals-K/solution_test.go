package leetcode

import "testing"

func TestSubarraySum(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 1, 1},
			k:        2,
			expected: 2,
		},
		{
			name:     "Example 2",
			nums:     []int{1, 2, 3},
			k:        3,
			expected: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := subarraySum(tc.nums, tc.k)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
