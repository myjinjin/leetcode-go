package leetcode

import (
	"reflect"
	"testing"
)

func TestPartitionLabels(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected []int
	}{
		{
			name:     "Example 1",
			s:        "ababcbacadefegdehijhklij",
			expected: []int{9, 7, 8},
		},
		{
			name:     "Example 2",
			s:        "eccbbbbdec",
			expected: []int{10},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := partitionLabels(tc.s)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
