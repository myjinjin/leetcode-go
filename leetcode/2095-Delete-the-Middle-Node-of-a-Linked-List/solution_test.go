package leetcode

import (
	"reflect"
	"testing"
)

func TestDeleteMiddle(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 3, 4, 7, 1, 2, 6},
			expected: []int{1, 3, 4, 1, 2, 6},
		},
		{
			input:    []int{1, 2, 3, 4},
			expected: []int{1, 2, 4},
		},
		{
			input:    []int{2, 1},
			expected: []int{2},
		},
	}

	for _, tc := range testCases {
		head := createLinkedList(tc.input)
		result := deleteMiddle(head)
		resultSlice := linkedListToSlice(result)
		if !reflect.DeepEqual(resultSlice, tc.expected) {
			t.Errorf("Input: %v, Expected: %v, Got: %v", tc.input, tc.expected, resultSlice)
		}
	}
}

func createLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}
	return head
}

func linkedListToSlice(head *ListNode) []int {
	result := []int{}
	current := head
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}
