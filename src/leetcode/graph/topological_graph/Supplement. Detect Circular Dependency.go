package topological_graph

func detectCircularDependency(edges [][]int) (bool, []int) {
	inDegreeList := make(map[int]int)
	for _, edge := range edges {
		if _, ok := inDegreeList[edge[0]]; !ok {
			inDegreeList[edge[0]] = 0
		}
		inDegreeList[edge[1]]++
	}

	var frontier []int
	// put the node with 0 in-degree into the frontier
	for node, inDegree := range inDegreeList {
		if inDegree == 0 {
			frontier = append(frontier, node)
		}
	}

	var path []int
	for len(frontier) > 0 {
		first := frontier[0]
		frontier = frontier[1:]
		for _, edge := range edges {
			if edge[0] != first {
				continue
			}
			inDegreeList[edge[1]]--
			if inDegreeList[edge[1]] == 0 {
				frontier = append(frontier, edge[1])
			}
		}
		path = append(path, first)
	}
	return len(path) == len(inDegreeList), path
}
