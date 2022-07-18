package find_if_path_exists_in_graph

func validPath(n int, edges [][]int, source int, destination int) bool {
	paths := make(map[int][]int)

	for _, e := range edges {
		paths[e[0]] = append(paths[e[0]], e[1])
		paths[e[1]] = append(paths[e[1]], e[0])
	}

	seen := make(map[int]bool)

	return traverse(source, destination, paths, seen)
}

func traverse(source int, destination int, paths map[int][]int, seen map[int]bool) bool {
	if source == destination {
		return true
	}

	if _, ok := seen[source]; ok {
		return false
	}
	seen[source] = true

	for _, path := range paths[source] {
		if v := traverse(path, destination, paths, seen); v {
			return true
		}
	}

	return false
}

// 10:54
// 11:04
// 10min
