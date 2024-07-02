package leetcode

import "testing"

func TestIsValid(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected bool
	}{
		{
			name:     "Example 1",
			s:        "()",
			expected: true,
		},
		{
			name:     "Example 2",
			s:        "()[]{}",
			expected: true,
		},
		{
			name:     "Example 3",
			s:        "(]",
			expected: false,
		},
		{
			name:     "Example 4",
			s:        "[",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isValid(tc.s)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
