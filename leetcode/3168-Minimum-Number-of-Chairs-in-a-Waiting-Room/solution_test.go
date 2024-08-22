package leetcode

import "testing"

func TestMinimumChairs(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "Example 1",
			s:        "EEEEEEE",
			expected: 7,
		},
		{
			name:     "Example 2",
			s:        "ELELEEL",
			expected: 2,
		},
		{
			name:     "Example 1",
			s:        "ELEELEELLL",
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := minimumChairs(tc.s)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
