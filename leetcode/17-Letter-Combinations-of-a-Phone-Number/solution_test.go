package leetcode

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	testCases := []struct {
		name     string
		digits   string
		expected []string
	}{
		{
			name:     "Example 1",
			digits:   "23",
			expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name:     "Example 2",
			digits:   "",
			expected: []string{},
		},
		{
			name:     "Example 3",
			digits:   "2",
			expected: []string{"a", "b", "c"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := letterCombinations(tc.digits)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}

// func letterCombinations(digits string) []string {

// }
