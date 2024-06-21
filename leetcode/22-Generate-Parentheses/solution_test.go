package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		expected []string
	}{
		{
			name:     "Example 1",
			n:        3,
			expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name:     "Example 2",
			n:        1,
			expected: []string{"()"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := generateParenthesis(tc.n)
			sort.Strings(result)
			sort.Strings(tc.expected)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
