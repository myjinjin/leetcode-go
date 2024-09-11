package leetcode

import (
	"testing"
)

func TestMinBitFlips(t *testing.T) {
	testCases := []struct {
		name     string
		start    int
		goal     int
		expected int
	}{
		{
			name:     "Example 1",
			start:    10,
			goal:     7,
			expected: 3,
		},
		{
			name:     "Example 2",
			start:    3,
			goal:     4,
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := minBitFlips(tc.start, tc.goal)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
