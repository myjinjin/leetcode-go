package leetcode

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	testCases := []struct {
		name     string
		list1    *ListNode
		list2    *ListNode
		expected *ListNode
	}{
		{
			name:     "Example 1",
			list1:    &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}},
			list2:    &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4}}}}}},
		},
		{
			name:     "Example 2",
			list1:    nil,
			list2:    nil,
			expected: nil,
		},
		{
			name:     "Example 3",
			list1:    nil,
			list2:    &ListNode{Val: 0},
			expected: &ListNode{Val: 0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := mergeTwoLists(tc.list1, tc.list2)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}
