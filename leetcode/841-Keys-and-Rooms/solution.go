package leetcode

func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	visitCount := 0
	var dfs func(room int)
	dfs = func(room int) {
		visited[room] = true
		visitCount++
		for _, key := range rooms[room] {
			if !visited[key] {
				dfs(key)
			}
		}
	}
	dfs(0)
	return visitCount == len(rooms)
}
