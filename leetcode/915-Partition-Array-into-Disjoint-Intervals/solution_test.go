package leetcode

import "testing"

func TestPartitionDisjoint(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{5, 0, 3, 8, 6},
			expected: 3,
		},
		{
			name:     "Example 2",
			nums:     []int{1, 1, 1, 0, 6, 12},
			expected: 4,
		},
		{
			name:     "Example 3",
			nums:     []int{1, 1},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := partitionDisjoint(tc.nums)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
