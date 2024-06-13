package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func TestCombinationSum3(t *testing.T) {
	testCases := []struct {
		name     string
		k        int
		n        int
		expected [][]int
	}{
		{
			name:     "Example 1",
			k:        3,
			n:        7,
			expected: [][]int{{1, 2, 4}},
		},
		{
			name:     "Example 2",
			k:        3,
			n:        9,
			expected: [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}},
		},
		{
			name:     "Example 3",
			k:        4,
			n:        1,
			expected: [][]int{},
		},
		{
			name:     "Example 4",
			k:        8,
			n:        36,
			expected: [][]int{{1, 2, 3, 4, 5, 6, 7, 8}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := combinationSum3(tc.k, tc.n)

			if !equalSlices(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}

func equalSlices(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		sort.Ints(a[i])
	}
	for i := range b {
		sort.Ints(b[i])
	}

	sort.Slice(a, func(i, j int) bool {
		return compare(a[i], a[j])
	})
	sort.Slice(b, func(i, j int) bool {
		return compare(b[i], b[j])
	})

	return reflect.DeepEqual(a, b)
}

func compare(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return false
}
