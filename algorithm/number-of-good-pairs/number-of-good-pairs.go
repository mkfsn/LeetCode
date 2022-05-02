package number_of_good_pairs

func numIdenticalPairs(nums []int) int {
	counts := make(map[int]int)

	for _, num := range nums {
		counts[num]++
	}

	count := 0

	for _, v := range counts {
		if v >= 2 {
			count += (v * (v - 1)) / 2
		}
	}

	return count
}

// 00:43
