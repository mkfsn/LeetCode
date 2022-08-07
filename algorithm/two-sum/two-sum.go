package two_sum

func twoSum(nums []int, target int) []int {
	cache := make(map[int][]int)

	for i, n := range nums {
		cache[n] = append(cache[n], i)
	}

	for i, n := range nums {
		indices, ok := cache[target-n]
		if !ok {
			continue
		}

		if len(indices) == 2 {
			return indices
		}

		if indices[0] == i {
			continue
		}

		return []int{i, indices[0]}
	}

	return []int{-1, -1}
}
