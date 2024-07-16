package leetcode

import "testing"

func TestLargestWordCount(t *testing.T) {
	testCases := []struct {
		name     string
		messages []string
		senders  []string
		expected string
	}{
		{
			name:     "Example 1",
			messages: []string{"Hello userTwooo", "Hi userThree", "Wonderful day Alice", "Nice day userThree"},
			senders:  []string{"Alice", "userTwo", "userThree", "Alice"},
			expected: "Alice",
		},
		{
			name:     "Example 2",
			messages: []string{"How is leetcode for everyone", "Leetcode is useful for practice"},
			senders:  []string{"Bob", "Charlie"},
			expected: "Charlie",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := largestWordCount(tc.messages, tc.senders)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
