package leetcode

type pair struct {
	node  string
	value float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	result := make([]float64, 0)

	graph := make(map[string][]pair)
	for i := 0; i < len(equations); i++ {
		equation := equations[i]
		start := equation[0]
		end := equation[1]
		value := values[i]
		graph[start] = append(graph[start], pair{end, value})
		graph[end] = append(graph[end], pair{start, 1 / value})
	}

	for _, query := range queries {
		start, end := query[0], query[1]
		if len(graph[start]) == 0 || len(graph[end]) == 0 {
			result = append(result, -1.0)
			continue
		}
		visited := make(map[string]bool)
		value := dfs(graph, visited, start, end, 1.0)
		result = append(result, value)
	}

	return result
}

func dfs(graph map[string][]pair, visited map[string]bool, start string, dest string, currVal float64) float64 {
	if visited[start] {
		return -1.0
	}
	visited[start] = true
	if start == dest {
		return currVal
	}
	pairs := graph[start]
	for _, pair := range pairs {
		if !visited[pair.node] {
			result := dfs(graph, visited, pair.node, dest, currVal*pair.value)
			if result != -1.0 {
				return result
			}
		}
	}
	visited[start] = false
	return -1.0
}
