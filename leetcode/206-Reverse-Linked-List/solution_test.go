package leetcode

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	testCases := []struct {
		name     string
		input    *ListNode
		expected *ListNode
	}{
		{
			name:     "Example 1",
			input:    createLinkedList([]int{1, 2, 3, 4, 5}),
			expected: createLinkedList([]int{5, 4, 3, 2, 1}),
		},
		{
			name:     "Example 2",
			input:    createLinkedList([]int{1, 2}),
			expected: createLinkedList([]int{2, 1}),
		},
		{
			name:     "Example 3",
			input:    createLinkedList([]int{}),
			expected: createLinkedList([]int{}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := reverseList(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

func createLinkedList(nums []int) *ListNode {
	var head *ListNode
	var current *ListNode
	for _, num := range nums {
		if head == nil {
			head = &ListNode{Val: num}
			current = head
		} else {
			current.Next = &ListNode{Val: num}
			current = current.Next
		}
	}
	return head
}
