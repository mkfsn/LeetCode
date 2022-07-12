package permutations

func permute(nums []int) [][]int {
	results := make([][]int, 0)
	paths := make([]int, len(nums))

	backtrack(0, len(nums), nums, paths, &results)

	return results
}

func backtrack(level int, N int, nums []int, paths []int, results *[][]int) {
	if level == N {
		res := make([]int, 0, N)
		for i := 0; i < N; i++ {
			res = append(res, nums[paths[i]])
		}
		*results = append(*results, res)
		return
	}

	for paths[level] = 0; paths[level] < N; paths[level]++ {
		ok := false
		for i := 0; i < level; i++ {
			if paths[level] == paths[i] {
				ok = true
			}
		}

		if ok {
			continue
		}

		backtrack(level+1, N, nums, paths, results)
	}
}
