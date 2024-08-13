package leetcode

import (
	"reflect"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	testCases := []struct {
		name       string
		candidates []int
		target     int
		expected   [][]int
	}{
		{
			name:       "Example 1",
			candidates: []int{10, 1, 2, 7, 6, 1, 5},
			target:     8,
			expected:   [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}},
		},
		{
			name:       "Example 2",
			candidates: []int{2, 5, 2, 1, 2},
			target:     5,
			expected:   [][]int{{1, 2, 2}, {5}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := combinationSum2(tc.candidates, tc.target)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
