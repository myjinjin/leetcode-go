package leetcode

import (
	"reflect"
	"testing"
)

func TestCommonChars(t *testing.T) {
	testCases := []struct {
		name     string
		words    []string
		expected []string
	}{
		{
			name:     "Example 1",
			words:    []string{"bella", "label", "roller"},
			expected: []string{"e", "l", "l"},
		},
		{
			name:     "Example 2",
			words:    []string{"cool", "lock", "cook"},
			expected: []string{"c", "o"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := commonChars(tc.words)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, Got: %v", tc.expected, result)
			}
		})
	}
}
