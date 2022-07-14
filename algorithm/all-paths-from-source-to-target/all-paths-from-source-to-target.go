package all_paths_from_source_to_target

func allPathsSourceTarget(graph [][]int) [][]int {
	res := make([][]int, 0)
	traverse([]int{0}, graph, &res)
	return res
}

func traverse(path []int, graph [][]int, res *[][]int) {
	cur := path[len(path)-1]

	if cur == len(graph)-1 {
		*res = append(*res, append([]int{}, path...))
		return
	}

	for _, step := range graph[cur] {
		newPath := append([]int{}, path...)
		traverse(append(newPath, step), graph, res)
	}
}
