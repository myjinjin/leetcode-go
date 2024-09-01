package leetcode

import (
	"reflect"
	"testing"
)

func TestConstruct2DArray(t *testing.T) {
	testCases := []struct {
		name     string
		original []int
		m        int
		n        int
		expected [][]int
	}{
		{
			name:     "Example 1",
			original: []int{1, 2, 3, 4},
			m:        2,
			n:        2,
			expected: [][]int{{1, 2}, {3, 4}},
		},
		{
			name:     "Example 2",
			original: []int{1, 2, 3},
			m:        1,
			n:        3,
			expected: [][]int{{1, 2, 3}},
		},
		{
			name:     "Example 3",
			original: []int{1, 2},
			m:        1,
			n:        1,
			expected: [][]int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := construct2DArray(tc.original, tc.m, tc.n)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
