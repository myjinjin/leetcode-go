package leetcode

import "testing"

func TestMinDistance(t *testing.T) {
	testCases := []struct {
		name     string
		word1    string
		word2    string
		expected int
	}{
		{
			name:     "Example 1",
			word1:    "horse",
			word2:    "ros",
			expected: 3,
		},
		{
			name:     "Example 2",
			word1:    "intention",
			word2:    "execution",
			expected: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := minDistance(tc.word1, tc.word2)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
