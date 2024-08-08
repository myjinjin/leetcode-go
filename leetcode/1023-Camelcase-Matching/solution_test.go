package leetcode

import (
	"reflect"
	"testing"
)

func TestCamelMatch(t *testing.T) {
	testCases := []struct {
		name     string
		queries  []string
		pattern  string
		expected []bool
	}{
		{
			name:     "Example 1",
			queries:  []string{},
			pattern:  "",
			expected: []bool{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := camelMatch(tc.queries, tc.pattern)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
