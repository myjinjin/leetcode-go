package leetcode

import "testing"

func TestCanVisitAllRooms(t *testing.T) {
	testCases := []struct {
		name     string
		rooms    [][]int
		expected bool
	}{
		{
			name:     "Example 1",
			rooms:    [][]int{{1}, {2}, {3}, {}},
			expected: true,
		},
		{
			name:     "Example 2",
			rooms:    [][]int{{1, 3}, {3, 0, 1}, {2}, {0}},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := canVisitAllRooms(tc.rooms)
			if result != tc.expected {
				t.Errorf("Expected: %v, Got: %v", tc.expected, result)
			}
		})
	}
}
