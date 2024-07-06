package leetcode

import "testing"

func TestSingleNonDuplicate(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 1, 2, 3, 3, 4, 4, 8, 8},
			expected: 2,
		},
		{
			name:     "Example 2",
			nums:     []int{3, 3, 7, 7, 10, 11, 11},
			expected: 10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := singleNonDuplicate(tc.nums)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
