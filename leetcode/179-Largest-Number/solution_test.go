package leetcode

import "testing"

func TestLargestNumber(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected string
	}{
		{
			name:     "Example 1",
			nums:     []int{10, 2},
			expected: "210",
		},
		{
			name:     "Example 2",
			nums:     []int{3, 30, 34, 5, 9},
			expected: "9534330",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := largestNumber(tc.nums)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
