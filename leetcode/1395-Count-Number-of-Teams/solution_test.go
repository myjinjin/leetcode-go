package leetcode

import (
	"testing"
)

func TestNumTeams(t *testing.T) {
	testCases := []struct {
		name     string
		rating   []int
		expected int
	}{
		{
			name:     "Example 1",
			rating:   []int{2, 5, 3, 4, 1},
			expected: 3,
		},
		{
			name:     "Example 2",
			rating:   []int{2, 1, 3},
			expected: 0,
		},
		{
			name:     "Example 3",
			rating:   []int{1, 2, 3, 4},
			expected: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := numTeams(tc.rating)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
