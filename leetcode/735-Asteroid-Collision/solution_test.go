package leetcode

import (
	"reflect"
	"testing"
)

func TestAsteroidCollision(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{5, 10, -5},
			expected: []int{5, 10},
		},
		{
			input:    []int{8, -8},
			expected: []int{},
		},
		{
			input:    []int{10, 2, -5},
			expected: []int{10},
		},
	}

	for _, tc := range testCases {
		result := asteroidCollision(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Input: %v, Expected: %v, Got: %v", tc.input, tc.expected, result)
		}
	}
}
