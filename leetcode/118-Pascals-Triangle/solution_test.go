package leetcode

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	testCases := []struct {
		name     string
		numRows  int
		expected [][]int
	}{
		{
			name:     "Example 1",
			numRows:  5,
			expected: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
		{
			name:     "Example 2",
			numRows:  1,
			expected: [][]int{{1}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := generate(tc.numRows)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
