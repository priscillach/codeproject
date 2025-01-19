package topological_graph

// https://leetcode.com/problems/course-schedule/description/
// finish times: 2
func canFinish(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		edges[prerequisites[i][1]] = append(edges[prerequisites[i][1]], prerequisites[i][0])
	}
	// visited
	// 0 never visited
	// mark 1 before visit next courses
	// mark 2 after visit all next courses of this course without a cycle
	visited := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		if !canFinishCore(i, edges, visited) {
			return false
		}
	}
	return true
}

func canFinishCore(course int, edges [][]int, visited []int) bool {
	if visited[course] == 2 {
		return true
	}
	if visited[course] == 1 {
		return false
	}
	visited[course] = 1
	for _, nextCourse := range edges[course] {
		if !canFinishCore(nextCourse, edges, visited) {
			return false
		}
	}
	visited[course] = 2
	return true
}
