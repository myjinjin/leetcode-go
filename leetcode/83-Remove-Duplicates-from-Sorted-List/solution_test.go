package leetcode

import "testing"

func TestDeleteDuplicates(t *testing.T) {
	testCases := []struct {
		name     string
		head     *ListNode
		expected *ListNode
	}{
		{
			name:     "Example 1",
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
		},
		{
			name:     "Example 2",
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}}},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
		},
		{
			name:     "Example 3",
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1}}},
			expected: &ListNode{Val: 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := deleteDuplicates(tc.head)
			if !equalList(result, tc.expected) {
				t.Errorf("Expected: %v, Got: %v", tc.expected, result)
			}
		})
	}
}

func equalList(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}
