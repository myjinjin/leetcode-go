package leetcode

import "testing"

func TestMaxUniqueSplit(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "Example 1",
			s:        "ababccc",
			expected: 5,
		},
		{
			name:     "Example 2",
			s:        "aba",
			expected: 2,
		},
		{
			name:     "Example 3",
			s:        "aa",
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maxUniqueSplit(tc.s)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
