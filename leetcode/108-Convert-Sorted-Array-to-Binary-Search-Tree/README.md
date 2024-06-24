# [108. Convert Sorted Array to Binary Search Tree](https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/)

## Problem

Given an integer array `nums` where the elements are sorted in **ascending order**, convert it to a **height-balanced binary search tree**.


Example 1:

![alt text](image.png)
![alt text](image-1.png)

```
Input: nums = [-10,-3,0,5,9]
Output: [0,-3,9,-10,null,5]
Explanation: [0,-10,5,null,-3,null,9] is also accepted:
```

Example 2:

![alt text](image-2.png)

```
Input: nums = [1,3]
Output: [3,1]
Explanation: [1,null,3] and [3,1] are both height-balanced BSTs.
```

Constraints:

- `1 <= nums.length <= 10^4`
- `-10^4 <= nums[i] <= 10^4`
- `nums` is sorted in a strictly increasing order.

## Solution

```go
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	n := len(nums)
	rootIndex := n / 2
	rootVal := nums[rootIndex]
	root := &TreeNode{Val: rootVal}
	left := sortedArrayToBST(nums[:rootIndex])
	right := sortedArrayToBST(nums[rootIndex+1:])
	root.Left = left
	root.Right = right
	return root
}
```