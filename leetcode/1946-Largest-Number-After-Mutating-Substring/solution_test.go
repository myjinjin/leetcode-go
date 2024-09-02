package leetcode

import "testing"

func TestMaximumNumber(t *testing.T) {
	testCases := []struct {
		name     string
		num      string
		change   []int
		expected string
	}{
		{
			name:     "Example 1",
			num:      "132",
			change:   []int{9, 8, 5, 0, 3, 6, 4, 2, 6, 8},
			expected: "832",
		},
		{
			name:     "Example 2",
			num:      "021",
			change:   []int{9, 4, 3, 5, 7, 2, 1, 9, 0, 6},
			expected: "934",
		},
		{
			name:     "Example 3",
			num:      "5",
			change:   []int{1, 4, 7, 5, 3, 2, 5, 6, 9, 4},
			expected: "5",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maximumNumber(tc.num, tc.change)
			if result != tc.expected {
				t.Errorf("Exepcted %v, but got %v", tc.expected, result)
			}
		})
	}
}
