package leetcode

import (
	"reflect"
	"testing"
)

func TestRelativeSortArray(t *testing.T) {
	testCases := []struct {
		name     string
		arr1     []int
		arr2     []int
		expected []int
	}{
		{
			name:     "Example 1",
			arr1:     []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
			arr2:     []int{2, 1, 4, 3, 9, 6},
			expected: []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
		},
		{
			name:     "Example 2",
			arr1:     []int{28, 6, 22, 8, 44, 17},
			arr2:     []int{22, 28, 8, 6},
			expected: []int{22, 28, 8, 6, 17, 44},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := relativeSortArray(tc.arr1, tc.arr2)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, Got: %v", tc.expected, result)
			}
		})
	}
}
