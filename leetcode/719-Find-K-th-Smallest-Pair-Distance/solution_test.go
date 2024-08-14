package leetcode

import "testing"

func TestSmallestDistancePair(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 3, 1},
			k:        1,
			expected: 0,
		},
		{
			name:     "Example 2",
			nums:     []int{1, 1, 1},
			k:        2,
			expected: 0,
		},
		{
			name:     "Example 3",
			nums:     []int{1, 6, 1},
			k:        3,
			expected: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := smallestDistancePair(tc.nums, tc.k)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
