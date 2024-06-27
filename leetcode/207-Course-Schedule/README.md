# [207. Course Schedule](https://leetcode.com/problems/course-schedule/)

## Problem

There are a total of `numCourses` courses you have to take, labeled from `0` to `numCourses - 1`. You are given an array `prerequisites` where `prerequisites[i] = [ai, bi]` indicates that you must take course bi first if you want to take course ai.

For example, the pair `[0, 1]`, indicates that to take course `0` you have to first take course `1`.
Return `true` if you can finish all courses. Otherwise, return `false`.


Example 1:

```
Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Explanation: There are a total of 2 courses to take. 
To take course 1 you should have finished course 0. So it is possible.
```

Example 2:

```
Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
Output: false
Explanation: There are a total of 2 courses to take. 
To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1. So it is impossible.
```

Constraints:

- `1 <= numCourses <= 2000`
- `0 <= prerequisites.length <= 5000`
- `prerequisites[i].length == 2`
- `0 <= ai, bi < numCourses`
- All the pairs `prerequisites[i]` are **unique**.

## Solution

```go
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	visited := make([]bool, numCourses)
	onStack := make([]bool, numCourses) // 현재 dfs 경로에 있는 노드 추적

	for _, pre := range prerequisites {
		course, preCourse := pre[0], pre[1]
		graph[course] = append(graph[course], preCourse)
	}

	var hasCycle func(course int) bool
	hasCycle = func(course int) bool {
        // 현재 dfs 경로에 이미 있는 과목이라면 사이클 존재
		if onStack[course] { 
			return true
		}

		if visited[course] {
			return false
		}

		visited[course] = true
		onStack[course] = true

        // 현재 노드의 선수과목들에 사이클 발견시 반환
		for _, pre := range graph[course] {
			if hasCycle(pre) {
				return true
			}
		}

		onStack[course] = false
		return false
	}

    // 모든 과목에 사이클 검사
	for course := 0; course < numCourses; course++ {
		if hasCycle(course) {
			return false
		}
	}
	return true
}
```