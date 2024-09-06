package leetcode

import (
	"reflect"
	"testing"
)

func TestModifiedList(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		head     *ListNode
		expected *ListNode
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3},
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			expected: &ListNode{Val: 4, Next: &ListNode{Val: 5}},
		},
		{
			name:     "Example 2",
			nums:     []int{1},
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}}}},
			expected: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2}}},
		},
		{
			name:     "Example 3",
			nums:     []int{5},
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := modifiedList(tc.nums, tc.head)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
