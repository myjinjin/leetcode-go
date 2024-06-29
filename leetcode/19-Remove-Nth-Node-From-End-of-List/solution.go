package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	sz := 0
	node := head
	for node != nil {
		sz++
		node = node.Next
	}

	if sz == n {
		head = head.Next
		return head
	}

	nth := head
	prev := &ListNode{}
	for i := 0; i < sz-n; i++ {
		if i == sz-n-1 {
			prev = nth
		}
		nth = nth.Next
	}

	prev.Next = nth.Next
	return head
}
