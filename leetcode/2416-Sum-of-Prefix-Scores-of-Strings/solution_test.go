package leetcode

import (
	"reflect"
	"testing"
)

func TestSumPrefixScores(t *testing.T) {
	testCases := []struct {
		name     string
		words    []string
		expected []int
	}{
		{
			name:     "Example 1",
			words:    []string{"abc", "ab", "bc", "b"},
			expected: []int{5, 4, 3, 2},
		},
		{
			name:     "Example 2",
			words:    []string{"abcd"},
			expected: []int{4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := sumPrefixScores(tc.words)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
