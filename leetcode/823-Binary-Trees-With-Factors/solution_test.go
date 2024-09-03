package leetcode

import "testing"

func TestNumFactoredBinaryTrees(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []int
		expected int
	}{
		{
			name:     "Example 1",
			arr:      []int{2, 4},
			expected: 3,
		},
		{
			name:     "Example 2",
			arr:      []int{2, 4, 5, 10},
			expected: 7,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := numFactoredBinaryTrees(tc.arr)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
