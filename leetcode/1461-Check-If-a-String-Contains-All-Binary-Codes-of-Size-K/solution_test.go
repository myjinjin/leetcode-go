package leetcode

import "testing"

func TestHasAllCodes(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		k        int
		expected bool
	}{
		{
			name:     "Example 1",
			s:        "00110110",
			k:        2,
			expected: true,
		},
		{
			name:     "Example 2",
			s:        "0110",
			k:        1,
			expected: true,
		},
		{
			name:     "Example 3",
			s:        "0110",
			k:        2,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := hasAllCodes(tc.s, tc.k)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
