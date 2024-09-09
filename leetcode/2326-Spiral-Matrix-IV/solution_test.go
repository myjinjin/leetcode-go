package leetcode

import (
	"reflect"
	"testing"
)

func TestSpiralMatrix(t *testing.T) {
	testCases := []struct {
		name     string
		m        int
		n        int
		head     *ListNode
		expected [][]int
	}{
		{
			name:     "Example 1",
			m:        3,
			n:        5,
			head:     &ListNode{Val: 3, Next: &ListNode{Val: 0, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 8, Next: &ListNode{Val: 1, Next: &ListNode{Val: 7, Next: &ListNode{Val: 9, Next: &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 5, Next: &ListNode{Val: 0}}}}}}}}}}}}},
			expected: [][]int{{3, 0, 2, 6, 8}, {5, 0, -1, -1, 1}, {5, 2, 4, 9, 7}},
		},
		{
			name:     "Example 2",
			m:        1,
			n:        4,
			head:     &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
			expected: [][]int{{0, 1, 2, -1}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := spiralMatrix(tc.m, tc.n, tc.head)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
