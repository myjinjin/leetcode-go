package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	result := &ListNode{}
	dummy := result
	sum := -1
	for head != nil {
		if head.Val == 0 {
			if sum != -1 {
				next := &ListNode{Val: sum}
				dummy.Next = next
				dummy = dummy.Next
			}
			sum = 0
		} else {
			sum += head.Val
		}
		head = head.Next
	}

	return result.Next
}
