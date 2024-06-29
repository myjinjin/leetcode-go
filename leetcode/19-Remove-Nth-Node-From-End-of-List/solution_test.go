package leetcode

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	testCases := []struct {
		name     string
		head     *ListNode
		n        int
		expected *ListNode
	}{
		{
			name:     "Example 1",
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			n:        2,
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}},
		},
		{
			name:     "Example 2",
			head:     &ListNode{Val: 1},
			n:        1,
			expected: nil,
		},
		{
			name:     "Example 3",
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			n:        1,
			expected: &ListNode{Val: 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := removeNthFromEnd(tc.head, tc.n)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
