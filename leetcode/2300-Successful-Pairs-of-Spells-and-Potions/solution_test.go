package leetcode

import (
	"reflect"
	"testing"
)

func TestSuccessfulPairs(t *testing.T) {
	testCases := []struct {
		name     string
		spells   []int
		potions  []int
		success  int64
		expected []int
	}{
		{
			name:     "Example 1",
			spells:   []int{5, 1, 3},
			potions:  []int{1, 2, 3, 4, 5},
			success:  7,
			expected: []int{4, 0, 3},
		},
		{
			name:     "Example 2",
			spells:   []int{3, 1, 2},
			potions:  []int{8, 5, 8},
			success:  16,
			expected: []int{2, 0, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := successfulPairs(tc.spells, tc.potions, tc.success)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
