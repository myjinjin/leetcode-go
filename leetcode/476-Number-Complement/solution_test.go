package leetcode

import "testing"

func TestFindComplement(t *testing.T) {
	testCases := []struct {
		name     string
		num      int
		expected int
	}{
		{
			name:     "Example 1",
			num:      5,
			expected: 2,
		},
		{
			name:     "Example 2",
			num:      1,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findComplement(tc.num)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
