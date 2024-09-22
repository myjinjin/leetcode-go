package leetcode

import "testing"

func TestFindKthNumber(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		k        int
		expected int
	}{
		{
			name:     "Example 1",
			n:        13,
			k:        2,
			expected: 10,
		},
		{
			name:     "Example 2",
			n:        1,
			k:        1,
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findKthNumber(tc.n, tc.k)
			if result != tc.expected {
				t.Errorf("Exepcted %v, but got %v", tc.expected, result)
			}
		})
	}
}
