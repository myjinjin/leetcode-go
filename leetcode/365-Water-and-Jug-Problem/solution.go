package leetcode

func canMeasureWater(x int, y int, target int) bool {
	if x+y < target {
		return false
	}

	visited := make(map[[2]int]bool)
	queue := [][2]int{{0, 0}}
	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		if state[0]+state[1] == target {
			return true
		}

		for _, nextState := range getNextStates(state, x, y) {
			if !visited[nextState] {
				visited[nextState] = true
				queue = append(queue, nextState)
			}
		}
	}
	return false
}

func getNextStates(state [2]int, x, y int) [][2]int {
	return [][2]int{
		{x, state[1]},
		{state[0], y},
		{0, state[1]},
		{state[0], 0},
		{min(state[0]+state[1], x), max(0, state[1]-(x-state[0]))},
		{max(0, state[0]-(y-state[1])), min(y, state[0]+state[1])},
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
