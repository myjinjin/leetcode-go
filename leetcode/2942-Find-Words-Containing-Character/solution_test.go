package leetcode

import (
	"reflect"
	"testing"
)

func TestFindWordsContaining(t *testing.T) {
	testCases := []struct {
		name     string
		words    []string
		x        byte
		expected []int
	}{
		{
			name:     "Example 1",
			words:    []string{"leet", "code"},
			x:        'e',
			expected: []int{0, 1},
		},
		{
			name:     "Example 2",
			words:    []string{"abc", "bcd", "aaaa", "cbc"},
			x:        'a',
			expected: []int{0, 2},
		},
		{
			name:     "Example 3",
			words:    []string{"abc", "bcd", "aaaa", "cbc"},
			x:        'z',
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findWordsContaining(tc.words, tc.x)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
