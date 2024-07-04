package leetcode

import "testing"

func TestMinDifference(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{5, 3, 2, 4},
			expected: 0,
		},
		{
			name:     "Example 2",
			nums:     []int{1, 5, 0, 10, 14},
			expected: 1,
		},
		{
			name:     "Example 3",
			nums:     []int{3, 100, 20},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := minDifference(tc.nums)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
