package leetcode

import (
	"reflect"
	"testing"
)

func TestOddEvenList(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 3, 5, 2, 4},
		},
		{
			input:    []int{2, 1, 3, 5, 6, 4, 7},
			expected: []int{2, 3, 6, 7, 1, 5, 4},
		},
	}

	for _, tc := range testCases {
		head := createLinkedList(tc.input)
		result := oddEvenList(head)
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
