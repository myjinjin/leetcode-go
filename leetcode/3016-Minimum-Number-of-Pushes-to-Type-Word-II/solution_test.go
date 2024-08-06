package leetcode

import "testing"

func TestMinimumPushes(t *testing.T) {
	testCases := []struct {
		name     string
		word     string
		expected int
	}{
		{
			name:     "Example 1",
			word:     "abcde",
			expected: 5,
		},
		{
			name:     "Example 2",
			word:     "xyzxyzxyzxyz",
			expected: 12,
		},
		{
			name:     "Example 3",
			word:     "aabbccddeeffgghhiiiiii",
			expected: 24,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := minimumPushes(tc.word)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
