package leetcode

import (
	"reflect"
	"testing"
)

func TestLexicalOrder(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		expected []int
	}{
		{
			name:     "Example 1",
			n:        13,
			expected: []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "Example 2",
			n:        2,
			expected: []int{1, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := lexicalOrder(tc.n)
			if !reflect.DeepEqual(tc.expected, result) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
