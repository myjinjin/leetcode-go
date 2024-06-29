package leetcode

import (
	"reflect"
	"testing"
)

func TestGetAncestors(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		edges    [][]int
		expected [][]int
	}{
		{
			name:     "Example 1",
			n:        8,
			edges:    [][]int{{0, 3}, {0, 4}, {1, 3}, {2, 4}, {2, 7}, {3, 5}, {3, 6}, {3, 7}, {4, 6}},
			expected: [][]int{nil, nil, nil, {0, 1}, {0, 2}, {0, 1, 3}, {0, 1, 2, 3, 4}, {0, 1, 2, 3}},
		},
		{
			name:     "Example 2",
			n:        5,
			edges:    [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
			expected: [][]int{nil, {0}, {0, 1}, {0, 1, 2}, {0, 1, 2, 3}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := getAncestors(tc.n, tc.edges)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
