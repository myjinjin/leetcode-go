package leetcode

func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	rowOrder := topologicalSort(k, rowConditions)
	if len(rowOrder) == 0 {
		return [][]int{}
	}

	colOrder := topologicalSort(k, colConditions)
	if len(colOrder) == 0 {
		return [][]int{}
	}

	rowIndex := make(map[int]int)
	colIndex := make(map[int]int)
	for i, val := range rowOrder {
		rowIndex[val] = i
	}
	for i, val := range colOrder {
		colIndex[val] = i
	}

	result := make([][]int, k)
	for i := range result {
		result[i] = make([]int, k)
	}

	for num := 1; num <= k; num++ {
		result[rowIndex[num]][colIndex[num]] = num
	}

	return result
}

func topologicalSort(k int, conditions [][]int) []int {
	graph := make(map[int][]int)
	indegree := make([]int, k+1)

	for _, c := range conditions {
		from, to := c[0], c[1]
		graph[from] = append(graph[from], to)
		indegree[to]++
	}

	var queue []int
	for i := 1; i <= k; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	var order []int
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		order = append(order, curr)

		for _, next := range graph[curr] {
			indegree[next]--
			if indegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	if len(order) != k {
		return []int{}
	}

	return order
}
