package leetcode

import "testing"

func TestLongestConsecutive(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{100, 4, 200, 1, 3, 2},
			expected: 4,
		},
		{
			name:     "Example 2",
			nums:     []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			expected: 9,
		},
		{
			name:     "Example 3",
			nums:     []int{0},
			expected: 1,
		},
		{
			name:     "Example 4",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "Example 5",
			nums:     []int{1, 2, 0, 1},
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := longestConsecutive(tc.nums)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
