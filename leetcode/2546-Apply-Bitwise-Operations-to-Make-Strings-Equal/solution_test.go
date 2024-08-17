package leetcode

import "testing"

func TestMakeStringsEqual(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		target   string
		expected bool
	}{
		{
			name:     "Example 1",
			s:        "1010",
			target:   "0110",
			expected: true,
		},
		{
			name:     "Example 2",
			s:        "11",
			target:   "00",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := makeStringsEqual(tc.s, tc.target)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
