package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	odd := head
	even := head.Next
	evenHead := even

	for even != nil && even.Next != nil {
		odd.Next = odd.Next.Next
		odd = odd.Next
		even.Next = even.Next.Next
		even = even.Next
	}

	odd.Next = evenHead

	return head
}
