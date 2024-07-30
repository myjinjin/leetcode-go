package leetcode

import "testing"

func TestMinimumDeletions(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "Example 1",
			s:        "aababbab",
			expected: 2,
		},
		{
			name:     "Example 2",
			s:        "bbaaaaabb",
			expected: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := minimumDeletions(tc.s)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
