package leetcode

import "testing"

func TestMaxWidthRamp(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{6, 0, 8, 2, 1, 5},
			expected: 4,
		},
		{
			name:     "Example 2",
			nums:     []int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1},
			expected: 7,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maxWidthRamp(tc.nums)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
