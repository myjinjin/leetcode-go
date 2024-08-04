package leetcode

import "testing"

func TestLargestPerimeter(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int64
	}{
		{
			name:     "Example 1",
			nums:     []int{5, 5, 5},
			expected: 15,
		},
		{
			name:     "Example 2",
			nums:     []int{1, 12, 1, 2, 5, 50, 3},
			expected: 12,
		},
		{
			name:     "Example 3",
			nums:     []int{5, 5, 50},
			expected: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := largestPerimeter(tc.nums)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
