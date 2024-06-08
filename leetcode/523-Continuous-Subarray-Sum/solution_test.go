package leetcode

import "testing"

func TestCheckSubarraySum(t *testing.T) {
	testCases := []struct {
		nums     []int
		k        int
		expected bool
	}{
		{
			nums:     []int{23, 2, 4, 6, 7},
			k:        6,
			expected: true,
		},
		{
			nums:     []int{23, 2, 6, 4, 7},
			k:        6,
			expected: true,
		},
		{
			nums:     []int{23, 2, 6, 4, 7},
			k:        13,
			expected: false,
		},
	}

	for _, tc := range testCases {
		result := checkSubarraySum(tc.nums, tc.k)
		if result != tc.expected {
			t.Errorf("Input: nums = %v, k = %d, Expected: %v, Got: %v", tc.nums, tc.k, tc.expected, result)
		}
	}
}
