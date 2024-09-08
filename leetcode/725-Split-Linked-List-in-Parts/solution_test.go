package leetcode

import (
	"reflect"
	"testing"
)

func TestSplitListToParts(t *testing.T) {
	testCases := []struct {
		name     string
		head     *ListNode
		k        int
		expected []*ListNode
	}{
		{
			name:     "Example 1",
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
			k:        5,
			expected: []*ListNode{{Val: 1}, {Val: 2}, {Val: 3}, nil, nil},
		},
		{
			name:     "Example 2",
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: &ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 10}}}}}}}}}},
			k:        3,
			expected: []*ListNode{{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}, {Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7}}}, {Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 10}}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := splitListToParts(tc.head, tc.k)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
