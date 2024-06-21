package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func TestSubsets(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3},
			expected: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			name:     "Example 2",
			nums:     []int{0},
			expected: [][]int{{}, {0}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := subsets(tc.nums)
			if !isEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}

func isEqual(a, b [][]int) bool {
	for i := range a {
		sort.Ints(a[i])
	}
	for i := range b {
		sort.Ints(b[i])
	}

	sort.Slice(a, func(i, j int) bool {
		return lessThan(a[i], a[j])
	})
	sort.Slice(b, func(i, j int) bool {
		return lessThan(b[i], b[j])
	})

	return reflect.DeepEqual(a, b)
}

func lessThan(a, b []int) bool {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}
