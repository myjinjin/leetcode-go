package leetcode

import "testing"

func TestFindMin(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{3, 4, 5, 1, 2},
			expected: 1,
		},
		{
			name:     "Example 2",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			expected: 0,
		},
		{
			name:     "Example 3",
			nums:     []int{11, 13, 15, 17},
			expected: 11,
		},
		{
			name:     "Example 4",
			nums:     []int{3, 1, 2},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findMin(tc.nums)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
