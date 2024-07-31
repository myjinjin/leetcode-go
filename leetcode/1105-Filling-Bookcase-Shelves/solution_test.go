package leetcode

import "testing"

func TestMinHeightShelves(t *testing.T) {
	testCases := []struct {
		name       string
		books      [][]int
		shelfWidth int
		expected   int
	}{
		{
			name:       "Example 1",
			books:      [][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}},
			shelfWidth: 4,
			expected:   6,
		},
		{
			name:       "Example 2",
			books:      [][]int{{1, 3}, {2, 4}, {3, 2}},
			shelfWidth: 6,
			expected:   4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := minHeightShelves(tc.books, tc.shelfWidth)
			if result != tc.expected {
				t.Errorf("Expectecd %v, but got %v", tc.expected, result)
			}
		})
	}
}
