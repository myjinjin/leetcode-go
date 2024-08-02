package leetcode

import (
	"reflect"
	"testing"
)

func TestDivisibilityArray(t *testing.T) {
	testCases := []struct {
		name     string
		word     string
		m        int
		expected []int
	}{
		{
			name:     "Example 1",
			word:     "998244353",
			m:        3,
			expected: []int{1, 1, 0, 0, 0, 1, 1, 0, 0},
		},
		{
			name:     "Example 2",
			word:     "1010",
			m:        10,
			expected: []int{0, 1, 0, 1},
		},
		{
			name:     "Example 3",
			word:     "91221181269244172125025075166510211202115152121212341281327",
			m:        21,
			expected: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := divisibilityArray(tc.word, tc.m)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
