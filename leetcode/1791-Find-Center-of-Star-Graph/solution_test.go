package leetcode

import "testing"

func TestFindCenter(t *testing.T) {
	testCases := []struct {
		name     string
		edges    [][]int
		expected int
	}{
		{
			name:     "Example 1",
			edges:    [][]int{{1, 2}, {2, 3}, {4, 2}},
			expected: 2,
		},
		{
			name:     "Example 2",
			edges:    [][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findCenter(tc.edges)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
