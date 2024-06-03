package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	temp := head
	size := 0
	for temp != nil {
		size++
		temp = temp.Next
	}

	prev := &ListNode{}
	prev.Next = head
	deleteNode := head
	middleIdx := size / 2
	currIdx := 0
	for deleteNode != nil && currIdx < middleIdx {
		currIdx++
		deleteNode = deleteNode.Next
		prev = prev.Next
	}

	prev.Next = deleteNode.Next
	return head
}
