package leetcode

import "testing"

func TestCanMeasureWater(t *testing.T) {
	testCases := []struct {
		name     string
		x        int
		y        int
		target   int
		expected bool
	}{
		{
			name:     "Example 1",
			x:        3,
			y:        5,
			target:   4,
			expected: true,
		},
		{
			name:     "Example 2",
			x:        2,
			y:        6,
			target:   5,
			expected: false,
		},
		{
			name:     "Example 3",
			x:        1,
			y:        2,
			target:   3,
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := canMeasureWater(tc.x, tc.y, tc.target)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
