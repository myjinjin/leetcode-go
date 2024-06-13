package leetcode

import "testing"

func TestPeakElement(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3, 1},
			expected: 2,
		},
		{
			name:     "Example 2",
			nums:     []int{1, 2, 1, 3, 5, 6, 4},
			expected: 5, //or 1
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findPeakElement(tc.nums)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
