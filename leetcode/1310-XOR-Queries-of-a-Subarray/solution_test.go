package leetcode

import (
	"reflect"
	"testing"
)

func TestXorQueries(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []int
		queries  [][]int
		expected []int
	}{
		{
			name:     "Example 1",
			arr:      []int{1, 3, 4, 8},
			queries:  [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}},
			expected: []int{2, 7, 14, 8},
		},
		{
			name:     "Example 2",
			arr:      []int{4, 8, 2, 10},
			queries:  [][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}},
			expected: []int{8, 0, 4, 4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := xorQueries(tc.arr, tc.queries)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
