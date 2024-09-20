package leetcode

import (
	"reflect"
	"testing"
)

func TestFindRepeatedDnaSequences(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected []string
	}{
		{
			name:     "Example 1",
			s:        "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT",
			expected: []string{"AAAAACCCCC", "CCCCCAAAAA"},
		},
		{
			name:     "Example 2",
			s:        "AAAAAAAAAAAAA",
			expected: []string{"AAAAAAAAAA"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findRepeatedDnaSequences(tc.s)
			if !reflect.DeepEqual(tc.expected, result) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
