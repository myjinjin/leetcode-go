package leetcode

import "testing"

func TestMaxSatisfied(t *testing.T) {
	testCases := []struct {
		name      string
		customers []int
		grumpy    []int
		minutes   int
		expected  int
	}{
		{
			name:      "Example 1",
			customers: []int{1, 0, 1, 2, 1, 1, 7, 5},
			grumpy:    []int{0, 1, 0, 1, 0, 1, 0, 1},
			minutes:   3,
			expected:  16,
		},
		{
			name:      "Example 2",
			customers: []int{1},
			grumpy:    []int{0},
			minutes:   1,
			expected:  1,
		},
		{
			name:      "Example 3",
			customers: []int{5, 8},
			grumpy:    []int{0, 1},
			minutes:   1,
			expected:  13,
		},
		{
			name:      "Example 4",
			customers: []int{3},
			grumpy:    []int{1},
			minutes:   1,
			expected:  3,
		},
		{
			name:      "Example 5",
			customers: []int{4, 10, 10},
			grumpy:    []int{1, 1, 0},
			minutes:   2,
			expected:  24,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maxSatisfied(tc.customers, tc.grumpy, tc.minutes)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
