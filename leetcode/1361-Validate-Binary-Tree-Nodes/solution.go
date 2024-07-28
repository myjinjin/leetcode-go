package leetcode

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = -1
	}

	for i := 0; i < n; i++ {
		if leftChild[i] != -1 {
			if parent[leftChild[i]] != -1 {
				return false
			}
			parent[leftChild[i]] = i
		}
		if rightChild[i] != -1 {
			if parent[rightChild[i]] != -1 {
				return false
			}
			parent[rightChild[i]] = i
		}
	}

	root := -1
	for i := 0; i < n; i++ {
		if parent[i] == -1 {
			if root != -1 {
				return false
			}
			root = i
		}
	}

	if root == -1 {
		return false
	}

	visited := make([]bool, n)
	var dfs func(node int)
	dfs = func(node int) {
		if node == -1 || visited[node] {
			return
		}
		visited[node] = true
		dfs(leftChild[node])
		dfs(rightChild[node])
	}

	dfs(root)

	for _, v := range visited {
		if !v {
			return false
		}
	}

	return true
}
