package topological_graph

// DFS
func findOrder(numCourses int, prerequisites [][]int) []int {
	edges := make([][]int, numCourses)
	visited := make([]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		edges[prerequisites[i][1]] = append(edges[prerequisites[i][1]], prerequisites[i][0])
	}
	var res []int
	for i := 0; i < numCourses; i++ {
		coursePath, canFinishFlag := findOrderCore(edges, visited, numCourses, i)
		if !canFinishFlag {
			return []int{}
		}
		res = append(coursePath, res...)
	}
	return res
}

func findOrderCore(edges [][]int, visited []int, numCourses int, course int) ([]int, bool) {
	if visited[course] == 1 {
		return []int{}, false
	}
	if visited[course] == 2 {
		return []int{}, true
	}
	visited[course] = 1
	var res []int
	for _, nextCourse := range edges[course] {
		coursePath, canFinishFlag := findOrderCore(edges, visited, numCourses, nextCourse)
		if !canFinishFlag {
			return []int{}, false
		}
		res = append(coursePath, res...)
	}
	// cur course the initial course
	res = append([]int{course}, res...)
	visited[course] = 2
	return res, true
}

// BFS todo
func findOrderV2(numCourses int, prerequisites [][]int) []int {
	var res []int
	return res
}
