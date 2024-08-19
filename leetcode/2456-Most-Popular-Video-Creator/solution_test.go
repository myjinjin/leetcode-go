package leetcode

import (
	"reflect"
	"testing"
)

func TestMostPopularCreator(t *testing.T) {
	testCases := []struct {
		name     string
		creators []string
		ids      []string
		views    []int
		expected [][]string
	}{
		{
			name:     "Example 1",
			creators: []string{"alice", "bob", "alice", "chris"},
			ids:      []string{"one", "two", "three", "four"},
			views:    []int{5, 10, 5, 4},
			expected: [][]string{{"alice", "one"}, {"bob", "two"}},
		},
		{
			name:     "Example 2",
			creators: []string{"alice", "alice", "alice"},
			ids:      []string{"a", "b", "c"},
			views:    []int{1, 2, 2},
			expected: [][]string{{"alice", "b"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := mostPopularCreator(tc.creators, tc.ids, tc.views)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
