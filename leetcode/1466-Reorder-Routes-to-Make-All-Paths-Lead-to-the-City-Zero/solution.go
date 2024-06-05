package leetcode

func minReorder(n int, connections [][]int) int {
	graph := make([][]int, n)
	for _, conn := range connections {
		start, end := conn[0], conn[1]
		graph[start] = append(graph[start], end)
		graph[end] = append(graph[end], -start)
	}
	visited := make([]bool, n)
	return dfs(graph, visited, 0)
}

func dfs(graph [][]int, visited []bool, start int) int {
	visited[start] = true
	count := 0
	for _, end := range graph[start] {
		if !visited[abs(end)] {
			if end > 0 {
				count++
			}
			count += dfs(graph, visited, abs(end))
		}
	}
	return count
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
