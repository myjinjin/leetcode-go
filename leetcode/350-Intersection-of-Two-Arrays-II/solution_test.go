package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func TestIntersect(t *testing.T) {
	testCases := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected []int
	}{
		{
			name:     "Example 1",
			nums1:    []int{1, 2, 2, 1},
			nums2:    []int{2, 2},
			expected: []int{2, 2},
		},
		{
			name:     "Example 2",
			nums1:    []int{4, 9, 5},
			nums2:    []int{9, 4, 9, 8, 4},
			expected: []int{4, 9},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := intersect(tc.nums1, tc.nums2)
			sort.Ints(result)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
