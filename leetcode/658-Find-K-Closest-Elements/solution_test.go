package leetcode

import (
	"reflect"
	"testing"
)

func TestFindClosestElements(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []int
		k        int
		x        int
		expected []int
	}{
		{
			name:     "Example 1",
			arr:      []int{1, 2, 3, 4, 5},
			k:        4,
			x:        3,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Example 2",
			arr:      []int{1, 2, 3, 4, 5},
			k:        4,
			x:        -1,
			expected: []int{1, 2, 3, 4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findClosestElements(tc.arr, tc.k, tc.x)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
