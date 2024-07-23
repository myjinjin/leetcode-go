package leetcode

import (
	"reflect"
	"testing"
)

func TestFrequencySort(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 1, 2, 2, 2, 3},
			expected: []int{3, 1, 1, 2, 2, 2},
		},
		{
			name:     "Example 2",
			nums:     []int{2, 3, 1, 3, 2},
			expected: []int{1, 3, 3, 2, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := frequencySort(tc.nums)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
