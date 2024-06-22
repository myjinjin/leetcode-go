package leetcode

import "testing"

func TestNumberOfSubarrays(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 1, 2, 1, 1},
			k:        3,
			expected: 2,
		},
		{
			name:     "Example 2",
			nums:     []int{2, 4, 6},
			k:        1,
			expected: 0,
		},
		{
			name:     "Example 3",
			nums:     []int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2},
			k:        2,
			expected: 16,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := numberOfSubarrays(tc.nums, tc.k)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
