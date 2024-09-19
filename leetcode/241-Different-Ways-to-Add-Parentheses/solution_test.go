package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func TestDiffWaysToCompute(t *testing.T) {
	testCases := []struct {
		name       string
		expression string
		expected   []int
	}{
		{
			name:       "Example 1",
			expression: "2-1-1",
			expected:   []int{0, 2},
		},
		{
			name:       "Example 2",
			expression: "2*3-4*5",
			expected:   []int{-34, -14, -10, -10, 10},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := diffWaysToCompute(tc.expression)
			sort.Ints(result)
			sort.Ints(tc.expected)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
