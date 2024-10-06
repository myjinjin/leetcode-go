package leetcode

import "testing"

func TestAreSentencesSimilar(t *testing.T) {
	testCases := []struct {
		name      string
		sentence1 string
		sentence2 string
		expected  bool
	}{
		{
			name:      "Example 1",
			sentence1: "My name is Haley",
			sentence2: "My Haley",
			expected:  true,
		},
		{
			name:      "Example 2",
			sentence1: "of",
			sentence2: "A lot of words",
			expected:  false,
		},
		{
			name:      "Example 3",
			sentence1: "Eating right now",
			sentence2: "Eating",
			expected:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := areSentencesSimilar(tc.sentence1, tc.sentence2)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
