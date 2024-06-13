package leetcode

import "testing"

func TestMinEatingSpeed(t *testing.T) {
	testCases := []struct {
		name     string
		piles    []int
		h        int
		expected int
	}{
		{
			name:     "Example 1",
			piles:    []int{3, 6, 7, 11},
			h:        8,
			expected: 4,
		},
		{
			name:     "Example 2",
			piles:    []int{30, 11, 23, 4, 20},
			h:        5,
			expected: 30,
		},
		{
			name:     "Example 3",
			piles:    []int{30, 11, 23, 4, 20},
			h:        6,
			expected: 23,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := minEatingSpeed(tc.piles, tc.h)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
	// piles []int, h int) int
}
