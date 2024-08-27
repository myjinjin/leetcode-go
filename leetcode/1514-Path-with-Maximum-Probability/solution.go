package leetcode

func maxProbability(n int, edges [][]int, succProb []float64, startNode int, endNode int) float64 {
	type neighborInfo struct {
		neighbor    int
		probability float64
	}

	graph := make(map[int][]neighborInfo)
	for i, edge := range edges {
		from, to := edge[0], edge[1]
		graph[from] = append(graph[from], neighborInfo{to, succProb[i]})
		graph[to] = append(graph[to], neighborInfo{from, succProb[i]})
	}

	maxProbabilities := make([]float64, n)
	maxProbabilities[startNode] = 1.0

	for i := 0; i < n-1; i++ {
		hasUpdate := false

		for currNode, neighbors := range graph {
			for _, neighborInfo := range neighbors {
				neighborNode := neighborInfo.neighbor
				edgeProbability := neighborInfo.probability

				if maxProbabilities[currNode]*edgeProbability > maxProbabilities[neighborNode] {
					maxProbabilities[neighborNode] = maxProbabilities[currNode] * edgeProbability
					hasUpdate = true
				}
			}
		}

		if !hasUpdate {
			break
		}
	}

	return maxProbabilities[endNode]
}
