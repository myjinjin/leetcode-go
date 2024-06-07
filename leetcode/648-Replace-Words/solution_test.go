package leetcode

import (
	"testing"
)

func TestReplaceWords(t *testing.T) {
	testCases := []struct {
		name       string
		dictionary []string
		sentence   string
		expected   string
	}{
		{
			name:       "Example 1",
			dictionary: []string{"cat", "bat", "rat"},
			sentence:   "the cattle was rattled by the battery",
			expected:   "the cat was rat by the bat",
		},
		{
			name:       "Example 2",
			dictionary: []string{"a", "b", "c"},
			sentence:   "aadsfasf absbs bbab cadsfafs",
			expected:   "a a b c",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := replaceWords(tc.dictionary, tc.sentence)
			if result != tc.expected {
				t.Errorf("Expected: %s, but got: %s", tc.expected, result)
			}
		})
	}
}
