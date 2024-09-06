package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	deleteMap := make(map[int]bool)
	for _, n := range nums {
		deleteMap[n] = true
	}

	temp := &ListNode{Next: head}
	curr := temp

	for curr.Next != nil {
		if _, exist := deleteMap[curr.Next.Val]; exist {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return temp.Next
}
