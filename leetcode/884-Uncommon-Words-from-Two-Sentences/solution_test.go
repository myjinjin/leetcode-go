package leetcode

import (
	"reflect"
	"testing"
)

func TestUncommonFromSentences(t *testing.T) {
	testCases := []struct {
		name     string
		s1       string
		s2       string
		expected []string
	}{
		{
			name:     "Example 1",
			s1:       "this apple is sweet",
			s2:       "this apple is sour",
			expected: []string{"sweet", "sour"},
		},
		{
			name:     "Example 2",
			s1:       "apple apple",
			s2:       "banana",
			expected: []string{"banana"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := uncommonFromSentences(tc.s1, tc.s2)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
