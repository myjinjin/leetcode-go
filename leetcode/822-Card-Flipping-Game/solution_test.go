package leetcode

import "testing"

func TestFlipgame(t *testing.T) {
	testCases := []struct {
		name     string
		fronts   []int
		backs    []int
		expected int
	}{
		{
			name:     "Example 1",
			fronts:   []int{1, 2, 4, 4, 7},
			backs:    []int{1, 3, 4, 1, 3},
			expected: 2,
		},
		{
			name:     "Example 2",
			fronts:   []int{1},
			backs:    []int{1},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := flipgame(tc.fronts, tc.backs)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
