package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	rightList := reverseList(slow.Next)

	leftList := head
	for rightList != nil {
		if leftList.Val != rightList.Val {
			return false
		}
		leftList = leftList.Next
		rightList = rightList.Next
	}

	return true
}

func reverseList(node *ListNode) *ListNode {
	var prev *ListNode
	curr := node
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
