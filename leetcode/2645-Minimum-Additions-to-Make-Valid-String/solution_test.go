package leetcode

import "testing"

func TestAddMinimum(t *testing.T) {
	testCases := []struct {
		name     string
		word     string
		expected int
	}{
		{
			name:     "Example 1",
			word:     "b",
			expected: 2,
		},
		{
			name:     "Example 2",
			word:     "aaa",
			expected: 6,
		},
		{
			name:     "Example 3",
			word:     "abc",
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := addMinimum(tc.word)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
