package leetcode

import (
	"reflect"
	"testing"
)

func TestCountBits(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		expected []int
	}{
		{
			name:     "Example 1",
			n:        2,
			expected: []int{0, 1, 1},
		},
		{
			name:     "Example 2",
			n:        5,
			expected: []int{0, 1, 1, 2, 1, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := countBits(tc.n)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
