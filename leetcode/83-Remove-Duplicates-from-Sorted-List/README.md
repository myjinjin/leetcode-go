# [83. Remove Duplicates from Sorted List](https://leetcode.com/problems/remove-duplicates-from-sorted-list/)

## Problem

Given the `head` of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

 

Example 1:

![alt text](image.png)

```
Input: head = [1,1,2]
Output: [1,2]
```

Example 2:

![alt text](image-1.png)

```
Input: head = [1,1,2,3,3]
Output: [1,2,3]
``` 

Constraints:

- The number of nodes in the list is in the range `[0, 300]`.
- `-100 <= Node.val <= 100`
- The list is guaranteed to be sorted in ascending order.

## Solution

```go
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
```