package leetcode

import "testing"

func TestHasCycle(t *testing.T) {
	testCases := []struct {
		name     string
		head     *ListNode
		expected bool
	}{
		{
			name:     "Example 1",
			head:     createLinkedList([]int{3, 2, 0, -4}, 1),
			expected: true,
		},
		{
			name:     "Example 2",
			head:     createLinkedList([]int{1, 2}, 0),
			expected: true,
		},
		{
			name:     "Example 3",
			head:     createLinkedList([]int{1}, -1),
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := hasCycle(tc.head)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
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
