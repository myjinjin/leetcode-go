package leetcode

import (
	"testing"
)

func TestRegionsBySlashes(t *testing.T) {
	testCases := []struct {
		name     string
		grid     []string
		expected int
	}{
		{
			name:     "Example 1",
			grid:     []string{" /", "/ "},
			expected: 2,
		},
		{
			name:     "Example 2",
			grid:     []string{" /", "  "},
			expected: 1,
		},
		{
			name:     "Example 3",
			grid:     []string{"/\\", "\\/"},
			expected: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := regionsBySlashes(tc.grid)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
