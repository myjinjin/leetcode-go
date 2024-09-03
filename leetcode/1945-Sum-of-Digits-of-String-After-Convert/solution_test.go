package leetcode

import (
	"testing"
)

func TestGetLucky(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		k        int
		expected int
	}{
		{
			name:     "Example 1",
			s:        "iiii",
			k:        1,
			expected: 36,
		},
		{
			name:     "Example 2",
			s:        "leetcode",
			k:        2,
			expected: 6,
		},
		{
			name:     "Example 3",
			s:        "zbax",
			k:        2,
			expected: 8,
		},
		{
			name:     "Example 4",
			s:        "fleyctuuajsr",
			k:        5,
			expected: 8,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := getLucky(tc.s, tc.k)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
