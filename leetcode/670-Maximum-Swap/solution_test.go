package leetcode

import "testing"

func TestMaximumSwap(t *testing.T) {
	testCases := []struct {
		name     string
		num      int
		expected int
	}{
		{
			name:     "Example 1",
			num:      2736,
			expected: 7236,
		},
		{
			name:     "Example 2",
			num:      9973,
			expected: 9973,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maximumSwap(tc.num)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
