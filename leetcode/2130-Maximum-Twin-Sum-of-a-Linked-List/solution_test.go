package leetcode

import "testing"

func TestPairSum(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{5, 4, 2, 1},
			expected: 6,
		},
		{
			input:    []int{4, 2, 2, 3},
			expected: 7,
		},
		{
			input:    []int{1, 100000},
			expected: 100001,
		},
	}

	for _, tc := range testCases {
		head := createLinkedList(tc.input)
		result := pairSum(head)
		if result != tc.expected {
			t.Errorf("Input: %v, Expected: %v, but got %v", tc.input, tc.expected, result)
		}
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
