package leetcode

import (
	"reflect"
	"testing"
)

func TestArrayRankTransform(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []int
		expected []int
	}{
		{
			name:     "Example 1",
			arr:      []int{},
			expected: []int{},
		},
		{
			name:     "Example 2",
			arr:      []int{},
			expected: []int{},
		},
		{
			name:     "Example 3",
			arr:      []int{},
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := arrayRankTransform(tc.arr)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Exepcted %v, but got %v", tc.expected, result)
			}
		})
	}
}
