package leetcode

import (
	"reflect"
	"testing"
)

func TestSortPeople(t *testing.T) {
	testCases := []struct {
		name     string
		names    []string
		heights  []int
		expected []string
	}{
		{
			name:     "Example 1",
			names:    []string{"Mary", "John", "Emma"},
			heights:  []int{180, 165, 170},
			expected: []string{"Mary", "Emma", "John"},
		},
		{
			name:     "Example 2",
			names:    []string{"Alice", "Bob", "Bob"},
			heights:  []int{155, 185, 150},
			expected: []string{"Bob", "Alice", "Bob"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := sortPeople(tc.names, tc.heights)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
