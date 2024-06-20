package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}

	prev := dummy
	curr := head

	for curr != nil {
		for curr.Next != nil && curr.Val == curr.Next.Val {
			curr = curr.Next
		}
		prev.Next = curr
		prev = prev.Next
		curr = curr.Next
	}
	return dummy.Next
}
