package leetcode

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	testCases := []struct {
		input    []byte
		expected []byte
	}{
		{
			input:    []byte{'h', 'e', 'l', 'l', 'o'},
			expected: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			input:    []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			expected: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}

	for _, tc := range testCases {
		reverseString(tc.input)
		if !reflect.DeepEqual(tc.input, tc.expected) {
			t.Errorf("Expected %v, but got %v", tc.expected, tc.input)
		}
	}
}
