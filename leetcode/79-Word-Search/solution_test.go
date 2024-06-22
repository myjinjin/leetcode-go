package leetcode

import "testing"

func TestExist(t *testing.T) {
	testCases := []struct {
		name     string
		board    [][]byte
		word     string
		expected bool
	}{
		{
			name:     "Example 1",
			board:    [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:     "ABCCED",
			expected: true,
		},
		{
			name:     "Example 2",
			board:    [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:     "SEE",
			expected: true,
		},
		{
			name:     "Example 3",
			board:    [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:     "ABCB",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := exist(tc.board, tc.word)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
