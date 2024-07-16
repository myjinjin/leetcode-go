# [2096. Step-By-Step Directions From a Binary Tree Node to Another](https://leetcode.com/problems/step-by-step-directions-from-a-binary-tree-node-to-another/)

## Problem

You are given the `root` of a binary tree with `n` nodes. Each node is uniquely assigned a value from `1` to `n`. You are also given an integer `startValue` representing the value of the start node `s`, and a different integer `destValue` representing the value of the destination node `t`.

Find the **shortest path** starting from node `s` and ending at node `t`. Generate step-by-step directions of such path as a string consisting of only the uppercase letters `'L'`, `'R'`, and `'U'`. Each letter indicates a specific direction:

- `'L'` means to go from a node to its left child node.
- `'R'` means to go from a node to its right child node.
- `'U'` means to go from a node to its parent node.

Return the step-by-step directions of the shortest path from node `s` to node `t`.


Example 1:

![alt text](image.png)

```
Input: root = [5,1,2,3,null,6,4], startValue = 3, destValue = 6
Output: "UURL"
Explanation: The shortest path is: 3 → 1 → 5 → 2 → 6.
```

Example 2:

![alt text](image-1.png)

```
Input: root = [2,1], startValue = 2, destValue = 1
Output: "L"
Explanation: The shortest path is: 2 → 1.
``` 

Constraints:

- The number of nodes in the tree is `n`.
- `2 <= n <= 10^5`
- `1 <= Node.val <= n`
- All the values in the tree are unique.
- `1 <= startValue, destValue <= n`
- `startValue != destValue`


## Solution

```go
func getDirections(root *TreeNode, startValue int, destValue int) string {
	lowestCommonAncestor := findLowestCommonAncestor(root, startValue, destValue)

	pathToStart := make([]byte, 0)
	pathToDest := make([]byte, 0)

	findPath(lowestCommonAncestor, &pathToStart, startValue)
	findPath(lowestCommonAncestor, &pathToDest, destValue)

	upwardPath := bytes.Repeat([]byte{'U'}, len(pathToStart))
	return string(upwardPath) + string(pathToDest)
}

func findLowestCommonAncestor(root *TreeNode, startValue, destValue int) *TreeNode {
	if root == nil || root.Val == startValue || root.Val == destValue {
		return root
	}

	left := findLowestCommonAncestor(root.Left, startValue, destValue)
	right := findLowestCommonAncestor(root.Right, startValue, destValue)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	return right
}

func findPath(node *TreeNode, path *[]byte, target int) bool {
	if node == nil {
		return false
	}

	if node.Val == target {
		return true
	}

	*path = append(*path, 'L')
	if findPath(node.Left, path, target) {
		return true
	}

	*path = (*path)[:len(*path)-1]

	*path = append(*path, 'R')
	if findPath(node.Right, path, target) {
		return true
	}

	*path = (*path)[:len(*path)-1]

	return false
}
```