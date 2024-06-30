package leetcode

import "testing"

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		name     string
		head     *ListNode
		expected bool
	}{
		{
			name:     "Example 1",
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}},
			expected: true,
		},
		{
			name:     "Example 2",
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isPalindrome(tc.head)
			if result != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
