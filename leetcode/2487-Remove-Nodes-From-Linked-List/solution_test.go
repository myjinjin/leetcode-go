package leetcode

import (
	"reflect"
	"testing"
)

func TestRemoveNodes(t *testing.T) {
	testCases := []struct {
		name     string
		head     *ListNode
		expected *ListNode
	}{
		{
			name:     "Example 1",
			head:     &ListNode{Val: 5, Next: &ListNode{Val: 2, Next: &ListNode{Val: 13, Next: &ListNode{Val: 3, Next: &ListNode{Val: 8}}}}},
			expected: &ListNode{Val: 13, Next: &ListNode{Val: 8}},
		},
		{
			name:     "Example 2",
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1}}}},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1}}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := removeNodes(tc.head)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
