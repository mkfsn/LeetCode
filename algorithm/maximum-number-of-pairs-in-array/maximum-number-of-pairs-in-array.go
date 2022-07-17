package maximum_number_of_pairs_in_array

func numberOfPairs(nums []int) []int {
	counts := make(map[int]int)

	for _, v := range nums {
		counts[v]++
	}

	pairs, left := 0, 0
	for _, cnt := range counts {
		pairs += cnt / 2
		left += cnt % 2
	}

	return []int{pairs, left}
}
