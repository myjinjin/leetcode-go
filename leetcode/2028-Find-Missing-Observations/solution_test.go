package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func TestMissingRolls(t *testing.T) {
	testCases := []struct {
		name     string
		rolls    []int
		mean     int
		n        int
		expected []int
	}{
		{
			name:     "Example 1",
			rolls:    []int{3, 2, 4, 3},
			mean:     4,
			n:        2,
			expected: []int{6, 6},
		},
		{
			name:     "Example 2",
			rolls:    []int{1, 5, 6},
			mean:     3,
			n:        4,
			expected: []int{2, 3, 2, 2},
		},
		{
			name:     "Example 3",
			rolls:    []int{1, 2, 3, 4},
			mean:     6,
			n:        4,
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := missingRolls(tc.rolls, tc.mean, tc.n)
			sort.Ints(result)
			sort.Ints(tc.expected)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
