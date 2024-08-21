package leetcode

import "testing"

func TestStrangePrinter(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "Example 1",
			s:        "aaabbb",
			expected: 2,
		},
		{
			name:     "Example 2",
			s:        "aba",
			expected: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := strangePrinter(tc.s)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
