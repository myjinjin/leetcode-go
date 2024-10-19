package leetcode

import "testing"

func TestFindKthBit(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		k        int
		expected byte
	}{
		{
			name:     "Example 1",
			n:        3,
			k:        1,
			expected: '0',
		},
		{
			name:     "Example 2",
			n:        4,
			k:        11,
			expected: '1',
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findKthBit(tc.n, tc.k)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
