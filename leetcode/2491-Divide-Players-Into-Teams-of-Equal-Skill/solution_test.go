package leetcode

import (
	"testing"
)

func TestDividePlayers(t *testing.T) {
	testCases := []struct {
		name      string
		skill     []int
		exepected int64
	}{
		{
			name:      "Example 1",
			skill:     []int{3, 2, 5, 1, 3, 4},
			exepected: 22,
		},
		{
			name:      "Example 2",
			skill:     []int{3, 4},
			exepected: 12,
		},
		{
			name:      "Example 3",
			skill:     []int{1, 1, 2, 3},
			exepected: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := dividePlayers(tc.skill)
			if result != tc.exepected {
				t.Errorf("Expected %v, but got %v", tc.exepected, result)
			}
		})
	}
}
