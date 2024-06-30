package leetcode

import (
	"reflect"
	"testing"
)

func TestDetectCycle(t *testing.T) {
	testCases := []struct {
		name     string
		head     *ListNode
		expected *ListNode
	}{
		{
			name:     "Example 1",
			head:     createLinkedList([]int{3, 2, 0, -4}, 1),
			expected: createLinkedList([]int{2, 0, -4}, 0),
		},
		{
			name:     "Example 2",
			head:     createLinkedList([]int{1, 2}, 0),
			expected: createLinkedList([]int{1, 2}, 0),
		},
		{
			name:     "Example 3",
			head:     createLinkedList([]int{1}, -1),
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := detectCycle(tc.head)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got %v", tc.expected, result)
			}
		})
	}
}

func createLinkedList(values []int, pos int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	curr := head
	var cycleNode *ListNode
	if pos == 0 {
		cycleNode = curr
	}

	for i := 1; i < len(values); i++ {
		curr.Next = &ListNode{Val: values[i]}
		curr = curr.Next
		if i == pos {
			cycleNode = curr
		}
	}

	if pos >= 0 && pos < len(values) {
		curr.Next = cycleNode
	}

	return head
}
