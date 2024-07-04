package leetcode

import (
	"testing"
)

func TestNextGreatestLetter(t *testing.T) {
	testCases := []struct {
		name     string
		letters  []byte
		target   byte
		expected byte
	}{
		{
			name:     "Example 1",
			letters:  []byte{'c', 'f', 'j'},
			target:   'a',
			expected: 'c',
		},
		{
			name:     "Example 2",
			letters:  []byte{'c', 'f', 'j'},
			target:   'c',
			expected: 'f',
		},
		{
			name:     "Example 3",
			letters:  []byte{'x', 'x', 'y', 'y'},
			target:   'z',
			expected: 'x',
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := nextGreatestLetter(tc.letters, tc.target)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
