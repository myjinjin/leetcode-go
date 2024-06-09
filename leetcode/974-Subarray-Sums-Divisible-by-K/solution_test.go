package leetcode

import "testing"

func TestSubarraysDivByK(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{4, 5, 0, -2, -3, 1},
			k:        5,
			expected: 7,
		},
		{
			name:     "Example 2",
			nums:     []int{5},
			k:        9,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := subarraysDivByK(tc.nums, tc.k)
			if result != tc.expected {
				t.Errorf("Expected: %v, Got: %v", tc.expected, result)
			}
		})
	}
}
