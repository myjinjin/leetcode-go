package leetcode

import (
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected [][]string
	}{
		{
			name:     "Example 1",
			s:        "aab",
			expected: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			name:     "Example 2",
			s:        "a",
			expected: [][]string{{"a"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := partition(tc.s)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
