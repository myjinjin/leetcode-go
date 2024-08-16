package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNodes(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	stack := []*ListNode{}
	curr := head

	for curr != nil {
		for len(stack) > 0 && stack[len(stack)-1].Val < curr.Val {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, curr)
		curr = curr.Next
	}

	for i := len(stack) - 1; i > 0; i-- {
		stack[i-1].Next = stack[i]
	}
	stack[len(stack)-1].Next = nil

	return stack[0]
}
