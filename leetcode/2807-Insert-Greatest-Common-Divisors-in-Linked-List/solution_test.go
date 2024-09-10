package leetcode

import (
	"reflect"
	"testing"
)

func TestInsertGreatestCommonDivisors(t *testing.T) {
	testCases := []struct {
		name     string
		head     *ListNode
		expected *ListNode
	}{
		{
			name:     "Example 1",
			head:     &ListNode{Val: 18, Next: &ListNode{Val: 6, Next: &ListNode{Val: 10, Next: &ListNode{Val: 3}}}},
			expected: &ListNode{Val: 18, Next: &ListNode{Val: 6, Next: &ListNode{Val: 6, Next: &ListNode{Val: 2, Next: &ListNode{Val: 10, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3}}}}}}},
		},
		{
			name:     "Example 2",
			head:     &ListNode{Val: 7},
			expected: &ListNode{Val: 7},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := insertGreatestCommonDivisors(tc.head)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
