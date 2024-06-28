package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	testCases := []struct {
		name     string
		strs     []string
		expected [][]string
	}{
		{
			name:     "Example 1",
			strs:     []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			name:     "Example 2",
			strs:     []string{""},
			expected: [][]string{{""}},
		},
		{
			name:     "Example 3",
			strs:     []string{"a"},
			expected: [][]string{{"a"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := groupAnagrams(tc.strs)
			if !compare(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}

func compare(expected, actual [][]string) bool {
	if len(expected) != len(actual) {
		return false
	}

	for i := range expected {
		sort.Strings(expected[i])
	}

	for i := range actual {
		sort.Strings(actual[i])
	}

	sort.Slice(expected, func(i, j int) bool {
		return less(expected[i], expected[j])
	})
	sort.Slice(actual, func(i, j int) bool {
		return less(actual[i], actual[j])
	})

	return reflect.DeepEqual(expected, actual)
}

func less(a, b []string) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}
