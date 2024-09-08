package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	n := 0
	curr := head
	for curr != nil {
		n++
		curr = curr.Next
	}

	baseSize := n / k
	extra := n % k

	result := make([]*ListNode, k)

	curr = head
	for i := 0; i < k; i++ {
		result[i] = curr

		size := baseSize
		if i < extra {
			size++
		}

		if curr != nil {
			for j := 1; j < size; j++ {
				curr = curr.Next
			}

			if curr != nil {
				next := curr.Next
				curr.Next = nil
				curr = next
			}
		}
	}

	return result
}
