package leetcode

import (
	"reflect"
	"testing"
)

func TestMergeNodes(t *testing.T) {
	testCases := []struct {
		name     string
		head     *ListNode
		expected *ListNode
	}{
		{
			name:     "Example 1",
			head:     &ListNode{Val: 0, Next: &ListNode{Val: 3, Next: &ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0}}}}}}}},
			expected: &ListNode{Val: 4, Next: &ListNode{Val: 11}},
		},
		{
			name:     "Example 2",
			head:     &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 3, Next: &ListNode{Val: 0, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0}}}}}}}},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := mergeNodes(tc.head)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
