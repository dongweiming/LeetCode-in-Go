package problem0207

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	visit := make([]int, numCourses)

	for _, pre := range prerequisites {
		graph[pre[0]] = append(graph[pre[0]], pre[1])
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(graph, visit, i) {
			return false
		}
	}
	return true
}

func dfs(graph [][]int, visit []int, i int) bool {
	if visit[i] == -1 {
		return false
	} else if visit[i] == 1 {
		return true
	}
	visit[i] = -1
	for _, j := range graph[i] {
		if !dfs(graph, visit, j) {
			return false
		}
	}
	visit[i] = 1
	return true
}
