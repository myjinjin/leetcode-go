package leetcode

func findCircleNum(isConnected [][]int) int {
	count := 0
	n := len(isConnected)
	visited := make([]bool, n)

	for start := 0; start < n; start++ {
		if !visited[start] {
			count += dfs(isConnected, visited, start)
		}
	}
	return count
}

func dfs(isConnected [][]int, visited []bool, start int) int {
	if visited[start] {
		return 0
	}
	count := 0
	visited[start] = true
	for dest := 0; dest < len(isConnected); dest++ {
		if isConnected[start][dest] == 1 && !visited[dest] {
			dfs(isConnected, visited, dest)
		}
	}
	count++
	return count
}
