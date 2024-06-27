package leetcode

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	visited := make([]bool, numCourses)
	onStack := make([]bool, numCourses)

	for _, pre := range prerequisites {
		course, preCourse := pre[0], pre[1]
		graph[course] = append(graph[course], preCourse)
	}

	var hasCycle func(course int) bool
	hasCycle = func(course int) bool {
		if onStack[course] {
			return true
		}

		if visited[course] {
			return false
		}

		visited[course] = true
		onStack[course] = true

		for _, pre := range graph[course] {
			if hasCycle(pre) {
				return true
			}
		}

		onStack[course] = false
		return false
	}

	for course := 0; course < numCourses; course++ {
		if hasCycle(course) {
			return false
		}
	}
	return true
}
