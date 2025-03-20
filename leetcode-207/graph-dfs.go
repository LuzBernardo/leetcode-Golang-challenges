package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)

	for _, prereq := range prerequisites {
		graph[prereq[0]] = append(graph[prereq[0]], prereq[1])
	}

	visited := make([]bool, numCourses)
	onPath := make([]bool, numCourses)

	var dfs func(int) bool
	dfs = func(course int) bool {
		if onPath[course] {
			return false
		}
		if visited[course] {
			return true
		}

		onPath[course] = true
		for _, prereq := range graph[course] {
			if !dfs(prereq) {
				return false
			}
		}

		onPath[course] = false
		visited[course] = true

		return true
	}

	for course := 0; course < numCourses; course++ {
		if !visited[course] && !dfs(course) {
			return false
		}
	}

	return true
}
