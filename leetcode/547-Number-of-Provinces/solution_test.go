package leetcode

import (
	"testing"
)

func TestFindCircleNum(t *testing.T) {
	testCases := []struct {
		name        string
		isConnected [][]int
		expected    int
	}{
		{
			name: "Example 1",
			isConnected: [][]int{
				{1, 1, 0},
				{1, 1, 0},
				{0, 0, 1},
			},
			expected: 2,
		},
		{
			name: "Example 2",
			isConnected: [][]int{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findCircleNum(tc.isConnected)
			if result != tc.expected {
				t.Errorf("Expected: %d, Got: %d", tc.expected, result)
			}
		})
	}
}
