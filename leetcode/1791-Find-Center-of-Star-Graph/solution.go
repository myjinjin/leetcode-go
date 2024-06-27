package leetcode

func findCenter(edges [][]int) int {
	nodes := make(map[int]bool)
	for _, edge := range edges {
		if _, ok := nodes[edge[0]]; ok {
			return edge[0]
		} else {
			nodes[edge[0]] = true
		}
		if _, ok := nodes[edge[1]]; ok {
			return edge[1]
		} else {
			nodes[edge[1]] = true
		}
	}
	return 0
}
